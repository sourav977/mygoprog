package main

import (
	"fmt"
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

type Todos struct {
	name  string
	lname string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}
	td1 := Todos{"sourav", "patnaik"}

	t, err := template.New("todos").Parse(`You have a task named "{{ .Name}}" with description: "{{ .Description}}"`)
	t1, _ := template.New("").Parse(`new task named "{{ .name}}" with description: "{{ .lname}}"`)
	t2 := template.Must(template.New("todos").Parse("You have task named \"{{ .Name}}\" with description: \"{{ .Description}}\""))
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	fmt.Println("")
	t1.Execute(os.Stdout, td1)
	fmt.Println("")
	t2.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}
