package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	req := readRequest(conn)
	writeResponse(conn, req)
}

func readRequest(conn net.Conn) (uri string) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fields := strings.Fields(ln)
		if i == 0 {
			fmt.Println("***METHOD", fields[0])
			fmt.Println("***URI", fields[1])
			uri = fields[1]
		}
		if i == 1 {
			fmt.Println("Hostname:", fields[1])
		}
		fmt.Println(ln)
		if ln == "" {
			break
		}
		i++
	}
	return uri
}

func writeResponse(conn net.Conn, req string) {
	content := "<!DOCTYPE html><html><header></header><body><p>This is an HTTP response</p></body></html>"
	if req != "/" {
		content = "<!DOCTYPE html><html><header></header><body><p>Page not found</p></body></html>"
	}
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, content)
}
