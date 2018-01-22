package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	funcMap := template.FuncMap{
		"flip": func(i int) string {
			if i == 0 {
				return "-0"
			}
			return fmt.Sprintf("%d", ^i+1)
		},
	}
	f, err := os.Create("lib.js")
	t := template.Must(template.New("template.txt").Funcs(funcMap).ParseFiles("template.txt"))
	err = t.Execute(f, make([]struct{}, 9007199254740992))
	if err != nil {
		log.Fatalln(err)
	}
}
