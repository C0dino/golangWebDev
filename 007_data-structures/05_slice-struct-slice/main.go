package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	f := car{
		Manufacturer: "Ford",
		Model:        "Focus",
		Doors:        2,
	}
	t := car{
		Manufacturer: "Tesla",
		Model:        "X",
		Doors:        4,
	}
	y := car{
		Manufacturer: "Tesla",
		Model:        "Y",
		Doors:        2,
	}
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}
	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with only love is healed",
	}

	sages := []sage{b, g, m}
	cars := []car{f, t, y}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
