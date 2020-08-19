package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {
		c, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	go handleInput(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Client:", ln)
	}
	defer conn.Close()

	fmt.Println("Client Disconnected")
}

func handleInput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}
