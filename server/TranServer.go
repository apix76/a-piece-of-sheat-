package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Trans struct {
	Path, Name string
	Send       net.Conn
}

func Sendfl(fileName string, conn net.Conn) {
	s := Trans{Path: os.Getenv("USERPROFILE") + "\\Downloads\\" + fileName, Send: conn}
	fileopen, err := os.Open(s.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileopen.Close()
	_, err = io.Copy(s.Send, fileopen)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(mes Message, send net.Conn) (fileName string) {
	fileName = time.Now().Format("20060102150405.000") + "_" + mes.File
	file, err := os.Create(os.Getenv("USERPROFILE") + "\\Downloads\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.CopyN(file, send, mes.FileSize)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}
