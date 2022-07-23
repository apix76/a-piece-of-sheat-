package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Text struct {
	Header, Text, Note string
	DateNote           time.Time
}

func CriateFile() {
	filePrt, err := os.Open("Text.txt")
	if err != nil {
		err := os.WriteFile("Text.txt", []byte("[]"), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer filePrt.Close()
}

func main() {
	var (
		m []Text
		r string
	)
zalopa:
	for {
		now := time.Now()
		CriateFile()
		fileRead, err := os.ReadFile("Text.txt")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(fileRead, &m)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Print("\033[2J\
		sort.Slice(m, func(i, j int) bool {
			if m[i].DateNote.IsZero() {
				return false
			}
			return m[i].DateNote.Before(m[j].DateNote)
		})
		for i, value := range m {
			fmt.Printf("[%d]. Заголовок: %s\nТекст: %s\nПримечание: %s\nDeadline: ", i, value.Header, value.Text, value.Note)
			if value.DateNote.IsZero() {
				fmt.Print("-\n\n")
			} else if value.DateNote.Before(now) {
				color.Red("%v\n\n", value.DateNote.Format("2006.01.02"))
			} else {
				fmt.Printf("%v\n\n", value.DateNote.Format("2006.01.02"))
			}
		}
		fmt.Println("Добавление[A] Удаление[D] Редактирование [R] Выход[E]")
		r = ScanLine()
		switch r {
		case "A", "a":
			m = Add(m)
		case "D", "d":
			m = Delet(m)
		case "R", "r":
			m = Replace(m)
		case "E", "e":
			break zalopa
		default:
			fmt.Println("Error")
		}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile("Text.txt", b, 0)
	}
}

func Delet(m []Text) []Text {
	var x int
	fmt.Print("Введите номер удаляемой записи: ")
	x, _ = strconv.Atoi(ScanLine())
	m = append(m[:x], m[x+1:]...)
	return m
}

func Add(m []Text) []Text {
	var a Text
	fmt.Print("Введите название заголовка: ")
	a.Header = ScanLine()
	fmt.Print("Введите текст: ")
	a.Text = ScanLine()
	fmt.Print("Введите заметку: ")
	a.Note = ScanLine()
	fmt.Print("Введите deadline или [-] : ")
	a.DateNote, _ = time.Parse("2006.01.02", ScanLine())
	m = append(m, a)
	return m
}

func Replace(m []Text) []Text {
	var (
		r string
		x int
	)
	for {
		fmt.Print("Введите номер изменяемой записи или [B] для выхода: ")
		if f := ScanLine(); f == "B" || f == "b" {
			return m
		} else {
			x, _ = strconv.Atoi(f)
		}
		fmt.Print("\nИзменить занголовок[H] Изменить текст[T] Изменить заметку[N] Изменить дедлайн [D] Назад[B]\n")
		r = ScanLine()
		switch r {
		case "H", "h":
			fmt.Print("Впишите изменение: ")
			m[x].Header = ScanLine()
		case "T", "t":
			fmt.Print("Впишите изменение: ")
			m[x].Text = ScanLine()
		case "N", "n":
			fmt.Print("Впишите изменение: ")
			m[x].Note = ScanLine()
		case "D", "d":
			fmt.Print("Впишите изменение: ")
			m[x].DateNote, _ = time.Parse("2006.01.02", ScanLine())
		case "B", "b":
			continue
		default:
			fmt.Print("Error")
		}
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
