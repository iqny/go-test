package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":10801")
	if err != nil {
		fmt.Println("listen:", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		client, err := l.Accept()
		if err != nil {
			fmt.Println("Accept", err)
			continue
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----", b)
	var host, port string
	switch b[3] {
	case 0x01: // IPv4
		host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		port = fmt.Sprintf("%d", b[8]<<8|b[9])
	case 0x03: // Domain name
		host = string(b[5 : n-2])
		port = fmt.Sprintf("%d", b[n-2]<<8|b[n-1])
	case 0x04: // IPv6
		host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		port = fmt.Sprintf("%d", b[20]<<8|b[21])
	}

	server, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x08, 0x43})
	go func() {
		_, err := io.Copy(server, client)
		if err != nil {
			fmt.Println("io.Copy error:", err)
		}
	}()
	_, err = io.Copy(client, server)
	if err != nil {
		fmt.Println("io.Copy error:", err)
	}
}
