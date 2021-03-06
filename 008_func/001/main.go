package main

import (
	"html/template"
	"strings"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

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

}
