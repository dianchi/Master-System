package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dianchi/Master-System/src"

	"github.com/kataras/iris/v12"
)

type Main struct {
	Name    string
	Mission string
}

func main() {
	dialogue := src.ReadDialogue()
	mission := src.ReadMission()
	b := Main{"Batrycc", mission}
	Web(ProcessString(dialogue, b))
}
func Web(message string) {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", message)
		ctx.View("hello.html")
		ctx.HTML("</br>Well 干得不错")
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
