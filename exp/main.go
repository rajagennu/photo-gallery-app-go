package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	t := template.Must(template.ParseFiles("hello.gohtml"))

	// fmt.Println(t)
	user := struct {
		Name     string
		Greeting string
	}{
		"Raja",
		"Morning",
	}

	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}
