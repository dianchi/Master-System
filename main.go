package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/dianchi/Master-System/src"

	"github.com/kataras/iris/v12"
)

type Main struct {
	Name    string
	Mission string
}

func main() {
	content, err := os.ReadFile("./Resources/Dialogue/Hello.json")
	if err != nil {
		fmt.Println(err)
	}
	dialogue := src.ReadDialogue()
	mission := src.ReadMission()
	options := src.ReadOptions(string(content))
	b := Main{"Batrycc", mission}
	Web(ProcessString(dialogue, b), options)
}
func Web(message string, options []string) {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", message)
		ctx.View("hello.html")
		for i := 0; i < len(options); i++ {
			ctx.HTML(options[i] + "</br>")
		}
	})

	app.Run(iris.Addr("localhost:8080"))
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		fmt.Println(err)
	}
	return tmplBytes.String()
}
func ProcessString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		fmt.Println(err)
	}
	return process(tmpl, vars)
}
