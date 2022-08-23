package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/toast.v1"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	up        = "\033[1F"
	clearline = "\033[2K"
	down      = "\033[1E"
)

/*

Чтобы получить файл от сервера, нужно отправить

*/

type Message struct {
	Text, Sender, Recipient string
	Time                    time.Time
	Request                 string // "download", "upload"
	File                    string
	FileSize                int64
}

var (
	chn = make(chan Message)
)

func main() {
	mes := Message{}
	send, err := net.Dial("tcp", "192.168.31.75:8080")
	if err != nil {
		log.Fatal(err)
	}
	go Serversend(send)
	go Gitgud(send)
	fmt.Print("Введите ваш никнейм: ")
	sen := ScanLine()
	fmt.Print("Введите никнейм получателя: ")
	rec := ScanLine()
	mes = Message{Sender: sen, Text: ""}
	Send(mes, send)
	for {
		fmt.Print("> ")
		text := ScanLine()
		fistword := strings.Split(text, " ")
		if fistword[0] == "/send" {
			s := Trans{Send: send}
			s.Path = fistword[1]
			s.Sendfl(sen, rec)
			continue
		}
		time := time.Now()
		mes = Message{Text: text, Time: time, Sender: sen, Recipient: rec}
		fmt.Print(up)
		fmt.Print(clearline)
		fmt.Printf("[%v] %v: %v\n", mes.Time.Format("15:04"), mes.Sender, mes.Text)
		fmt.Print(down)
		chn <- mes
	}
}

func ScanLine() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода: ", err)
	}
	return in.Text()
}

func Serversend(send net.Conn) {
	for {
		stec, ok := <-chn
		if ok == false {
			break
		}
		Send(stec, send)
	}
}

func Gitgud(send net.Conn) {
	mes := Message{}
	scan := bufio.NewScanner(send)
	for scan.Scan() {
		fmt.Print("\r")
		err := json.Unmarshal(scan.Bytes(), &mes)
		if err != nil {
			log.Fatal(err)
		}
		fileName := strings.SplitN(mes.File, "_", 2)
		if mes.File != "" {
			size, lenth := Filelenth(int(mes.FileSize))
			fmt.Printf("Вам отправили файл %v %v %v, скачать? [Y/N]", fileName[1], size, lenth)
			a := ScanLine()
			switch a {
			case "y", "Y":
				s := Trans{Send: send}
				s.Name = mes.File
				s.Get(mes.Sender, mes.Recipient, mes, int(mes.FileSize), fileName)
			default:
			}
			continue
		}
		fmt.Printf("[%v] %v: %v\n", mes.Time.Format("15:04"), mes.Sender, mes.Text)
		fmt.Print("> ")
		notify := toast.Notification{
			AppID:   "nikita_chat",
			Title:   "Message from " + mes.Sender,
			Message: mes.Text,
		}
		if err := notify.Push(); err != nil {
			log.Fatal(err)
		}
	}
}

func Filelenth(size int) (int, string) {
	var lenth string
	count := 0
	for size/1024 != 0 {
		count += 1
		size /= 1024
	}
	switch count {
	case 0:
		lenth = "byte"
	case 1:
		lenth = "Kb"
	case 2:
		lenth = "Mb"
	case 3:
		lenth = "Gb"

	}
	return size, lenth
}

func Send(mes Message, con net.Conn) {
	bitstec, err := json.Marshal(mes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = con.Write(bitstec)
	if err != nil {
		log.Fatal(err)
	}
	_, err = con.Write([]byte("\n"))
	if err != nil {
		log.Fatal(err)
	}
}

//s := sync.Mutex{}
//s.Lock()
//s.Unlock()
//c := make(chan Message)
//go func() { c <- Message{"Text", Time.Now(), "aaa", "bbbb"} }()
//msg := <-c
//fmt.Println(msg.Text, msg.Recipient, msg.Sender, msg.Time.Format("15:04"))
