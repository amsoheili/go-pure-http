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
		fmt.Println(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	uri := request(conn)
	fmt.Println(uri)
	switch uri {
	case "/talk":
		respond(conn, "hello babe")
	case "/sing":
		respond(conn, "I'm gonna take my horse to the old town road ...")
	default:
		respond(conn, "I don't know the route")
	}
}

func request(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	requestedUri := ""
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			requestedUri = strings.Fields(line)[1]
		}
		if line == "" {
			break
		}
		i++
	}
	return requestedUri
}

func respond(conn net.Conn, bodyStr string) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
<div>` + bodyStr + `</div></body></html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
	conn.Close()
}
