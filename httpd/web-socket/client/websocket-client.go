package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8090")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	defer conn.Close()
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
