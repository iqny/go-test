package main

import (
	"fmt"
)

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type WordsCounter int

/*func (w *WordsCounter) Write(p []byte) (n int, err error) {
	s:=strings.NewReader(string(p))
	bs:=bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	sum:=0
	for bs.Scan(){

	}
}*/

