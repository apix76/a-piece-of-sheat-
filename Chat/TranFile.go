package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

type Trans struct {
	Path, Name string
	Send       net.Conn
}

func (s *Trans) Sendfl(sen string, res string) {
	file, err := os.Stat(s.Path)
	if err != nil {
		log.Fatal(err)
	}
	filesize := file.Size()
	if err != nil {
		log.Fatal(err)
	}
	_, filename := filepath.Split(s.Path)
	mes := Message{Request: "upload",
		FileSize: filesize,
		File:     filename, Sender: sen, Recipient: res}
	Send(mes, s.Send)
	fileopen, err := os.Open(s.Path)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(s.Send, fileopen)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Trans) Get(mes Message, Filesize int, fileName []string) {
	mesget := Message{Request: "download", File: s.Name, FileSize: int64(Filesize)}
	Send(mesget, s.Send)
	file, err := os.Create(os.Getenv("USERPROFILE") + "\\Downloads\\" + fileName[1])
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.CopyN(file, s.Send, mes.FileSize)
	if err != nil {
		log.Fatal(err)
	}
}
