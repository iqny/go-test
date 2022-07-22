package gridge

import (
	"fmt"
	"testing"
)

func TestBridgeHtml(t *testing.T) {
	formatter :=New(HtmlFormatter{})
	fmt.Println(formatter.get())
	formatter.setFormatter(PlainTextFormatter{})
	fmt.Println(formatter.get())
}
func TestBridgePinText(t *testing.T) {
	formatter:=New(PlainTextFormatter{})
	fmt.Println(formatter.get())
}
