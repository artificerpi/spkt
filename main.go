package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

var (
	server   string
	port     int
	protocol string
	filename string
)

func main() {
	// parse flags
	flag.StringVar(&server, "server", "localhost", "specify the server address to send packet")
	flag.IntVar(&port, "port", 5678, "specify the port to send packet")
	flag.StringVar(&protocol, "proto", "tcp", "specify the protocol (tcp or udp)")
	flag.StringVar(&filename, "file", "", "read data to send from file")
	flag.Parse()

	args := flag.Args()
	// todo input validation

	// set data
	var buf []byte
	var err error
	if len(args) == 0 {
		// default greeting for testing
		buf = []byte("Hello there!")
	} else {
		// take space as subsequent input
		buf = []byte(strings.Join(args, ""))
	}

	if filename != "" {
		buf, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	addr := fmt.Sprintf("%s:%d", server, port)
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// send packet
	if sendMsg(conn, buf) {
		fmt.Println("Send message successfully.")
	} else {
		fmt.Println("Failed to send message.")
	}
}

func sendMsg(conn net.Conn, buf []byte) bool {
	n, err := conn.Write(buf)
	if n != len(buf) || err != nil {
		log.Println(err)
		return false
	}

	return true
}
