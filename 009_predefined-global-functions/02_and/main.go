package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}
	u2 := user{
		Name:  "Ghandi",
		Motto: "Be the change",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
