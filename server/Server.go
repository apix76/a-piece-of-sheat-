package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

type Message struct {
	Text, Sender, Recipient string
	Time                    time.Time
	Request                 string // "download", "upload"
	File                    string
	FileSize                int64
}

var (
	ClientConn = make(map[string]net.Conn)
)

func main() {
	listed, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	Acept(listed)
}

func Acept(listed net.Listener) {
	for {
		c, err := listed.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Handle(c)
	}
}

func Handle(Conn net.Conn) {
	var (
		mes      Message
		fileName string
	)
	defer Conn.Close()
	scan := bufio.NewScanner(Conn)
	for scan.Scan() {
		err := json.Unmarshal(scan.Bytes(), &mes)
		if err != nil {
			fmt.Printf("Какае-то хуйня, разберись! \nerr: %v \nip: %v\n", err, Conn.RemoteAddr())
			return
		}
		if mes.Request == "upload" {
			fileName = Get(mes, ClientConn[mes.Sender])
			mes = Message{File: fileName}
		}
		if mes.Request == "download" {
			Sendfl(fileName, ClientConn[mes.Recipient])
		}
		if ClientConn[mes.Sender] == nil {
			ClientConn[mes.Sender] = Conn
		}
		if ClientConn[mes.Recipient] != nil {
			err = Send(mes, ClientConn[mes.Recipient])
			if err != nil {
				EndOfChat(mes.Sender, Conn)
				delete(ClientConn, mes.Recipient)
				break
			}
		}
	}
	delete(ClientConn, mes.Sender)
	if ClientConn[mes.Recipient] != nil {
		EndOfChat(mes.Recipient, ClientConn[mes.Recipient])
	}
}

func Send(mes Message, con net.Conn) (err error) {
	bitstec, err := json.Marshal(mes)
	if err != nil {
		return err
	}
	_, err = con.Write(bitstec)
	if err != nil {
		return err
	}
	_, err = con.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}

func EndOfChat(m string, Conn net.Conn) {
	off := Message{Sender: "SERVER", Recipient: m, Text: "Собеседник поинул диалог."}
	err := Send(off, Conn)
	if err != nil {
		Conn.Close()
	}
}
