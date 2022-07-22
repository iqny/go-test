package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	for {
		result := make([]byte, 256)
		n, err := conn.Read(result)
		if err != nil {
			break
		}
		fmt.Printf("%s", result)
		if 0 == n {
			break
		} else if strings.TrimSpace(string(result[:n])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
