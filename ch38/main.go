package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"log"
	"os"
	"time"
)

// {{ .xxoo -}} 删除右侧的空白
var md = `个人信息:
姓名: {{ .Name }}
年龄: {{ .Age }}
爱好: {{ .Hobby -}}
`

type People struct {
	Name string
	Age  int
}

func (p People) Hobby() string {
	return "唱,跳,rap,篮球"
}
func main() {
	tpl := template.Must(template.New("first").Parse(md))
	p := People{
		Name: "Jackson",
		Age:  20,
	}
	py, _ := json.Marshal(p)
	fmt.Println(string(py))
	if err := tpl.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
	uuidByte, _ := uuid.NewUUID()
	fmt.Println(fmt.Sprintf("%s", uuidByte.String()))
	var ii int64 = 999
	fmt.Println(time.Duration(ii))
}
