package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data any) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	// Execute: Write out the template text
	// os.Stdout: Write the template to the terminal
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	text := "Here's my template!\n"
	executeTemplate(text, nil)

	fmt.Println()

	// To insert data in a template, we need to add 'actions' to the template text.
	// Actions are denoted with double curly braces, {{ }}. Inside it, you specify data you want to insert.
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	executeTemplate(templateText, "ABC")
	executeTemplate(templateText, 42)
	executeTemplate(templateText, true)
	// We can use a single period to refer do 'dot', the current value within the data the template is working with.
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.45)
}
