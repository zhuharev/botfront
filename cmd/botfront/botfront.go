package main

import (
	"log"

	"github.com/zhuharev/botfront/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalln(err)
	}
	srv.Run()
}
