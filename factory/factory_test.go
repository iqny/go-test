package factory

import (
	"testing"
)

func TestSimpleParseFactory_Create(t *testing.T) {
	simple:=SimpleParseFactory{}
	parse:=simple.Create("yaml")
	if parse == nil {
		t.Fatal("parse is nil")
	}
	data:=parse.Parse("config/config.yaml")
	t.Log(data)
}

func TestCreateFactory(t *testing.T) {
	factory:=CreateFactory("json")
	parse :=factory.createParse()
	if parse!=nil{
		data := parse.Parse("config/config.json")
		t.Log(data)
	}
}
