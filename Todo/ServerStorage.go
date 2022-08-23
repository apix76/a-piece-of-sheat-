package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ServerStorage struct {
	listID string
}

func (s *ServerStorage) Replace(x int, upd TextUpdate) {
	method := "PATCH"
	url := fmt.Sprintf("http://192.168.31.202:8080/todo/%s/%d", s.listID, x)
	updBody, _ := json.Marshal(upd)
	updBodyR := bytes.NewReader(updBody)
	req, _ := http.NewRequest(method, url, updBodyR)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode/100 != 2 {
		log.Fatal("Error: ", resp.StatusCode)
	}
}

func (s *ServerStorage) Add(upd Text) {
	method := "POST"
	url := fmt.Sprintf("http://192.168.31.202:8080/todo/%s", s.listID)
	updBody, _ := json.Marshal(upd)
	updBodyR := bytes.NewReader(updBody)
	req, _ := http.NewRequest(method, url, updBodyR)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode/100 != 2 {
		log.Fatal("Error: ", resp.StatusCode)
	}
}

func (s *ServerStorage) Delet(x int) {
	method := "DELETE"
	url := fmt.Sprintf("http://192.168.31.202:8080/todo/%s/%d", s.listID, x)
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode/100 != 2 {
		log.Fatal("Error: ", resp.StatusCode)
	}
}

func (s *ServerStorage) Get() []Text {
	var (
		bit []byte
		m   []Text
	)
	method := "GET"
	url := fmt.Sprintf("http://192.168.31.202:8080/todo/%s", s.listID)
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode/100 != 2 {
		log.Fatal("Error: ", resp.StatusCode)
	}
	bit, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	err = json.Unmarshal(bit, &m)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return m
}
