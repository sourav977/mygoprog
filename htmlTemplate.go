package main

import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	// Parse data -- omitted for brevity
	td := ToDo{User: "sourav", List: []entry{{Name: "sikan", Done: true}, {Name: "tej", Done: false}}}

	// Files are provided as a slice of strings.
	paths := []string{
		"todo.tmpl",
	}

	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}
