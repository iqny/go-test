package main

import (
	"fmt"
	"net/url"
)

func main() {
	m:=url.Values{"lang":{"cn"}}
	m.Add("item","1")
	m.Add("item","")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("item"))
	fmt.Println(m.Get("q"))

	m = nil
	fmt.Println(m.Get("o"))
	fmt.Printf("bbb%s",m)
	array1:=make(map[string]string)
	array2:=make(map[string]string)
	array1["a"] = "1"
	array2["b"] = "2"
	merge(array1,array2)
	fmt.Println()
	fmt.Println(array1)
	fmt.Println(21212)
}

func merge(array1,arrar2 map[string]string)map[string]string  {
	newArrar := make(map[string]string,len(array1)+len(arrar2))
	for i,v:=range array1 {
		newArrar[i] = v
	}
	for i,v:=range arrar2 {
		newArrar[i] = v
	}
	return newArrar
}