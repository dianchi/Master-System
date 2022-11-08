package main

import (
	"bytes"
	"fmt"

	//"fmt"
	"os"
	"text/template"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

type Main struct {
	Name    string
	Mission string
}

func main() {
	dialogue := ReadDialogue()
	mission := ReadMission()
	b := Main{"Batrycc", mission}
	Web(ProcessString(dialogue, b))
}
func Web(message string) {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", message)
		ctx.View("hello.html")
	})

	app.Run(iris.Addr("localhost:8080"))
}
func ReadDialogue() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dialogue := viper.New()
	dialogue.AddConfigPath(path + "/Resources/Dialogue")
	dialogue.SetConfigName("CTM-1")
	dialogue.SetConfigType("yaml")
	if err := dialogue.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return dialogue.GetString("Contents")
}
func ReadMission() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	mission := viper.New()
	mission.AddConfigPath(path + "/Resources/Missions")
	mission.SetConfigName("CTM-Mission")
	mission.SetConfigType("yaml")
	if err := mission.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return mission.GetString("Contents")
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
