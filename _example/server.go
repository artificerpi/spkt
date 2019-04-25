package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":5678")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("I'm listenning at :5678")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go recvMsg(conn)
	}
}

func recvMsg(conn net.Conn) {
	defer conn.Close()

	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(data))
}
