package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host> <port>\n", os.Args[0])
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", addr)

	go readFromServer(conn)
	writeToServer(conn)
}

func readFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from server: %v\n", err)
	}
}

func writeToServer(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintf(conn, "%s\n", text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to server: %v\n", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
	}
}

// go run main.go <host> <port>
// go run main.go localhost 8080
