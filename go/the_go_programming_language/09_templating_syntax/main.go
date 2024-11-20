package main

import (
    "os"
    "text/template"
)

type Person struct {
    Name    string
    Age     int
    Hobbies []string
}

func main() {
    // Define a template string
    tmpl := `
    Hello, {{.Name}}!
    {{if .Age}}You are {{.Age}} years old.{{else}}Your age is unknown.{{end}}
    
    {{if .Hobbies}}
    Your hobbies are:
    {{range .Hobbies}}- {{.}}
    {{end}}
    {{else}}
    You don't have any hobbies listed.
    {{end}}
    `

    // Parse the template
    t, err := template.New("example").Parse(tmpl)
    if err != nil {
        panic(err)
    }

    // Execute the template with a sample data
    person := Person{
        Name:    "Alice",
        Age:     30,
        Hobbies: []string{"Reading", "Hiking", "Cooking"},
    }
    t.Execute(os.Stdout, person)
}
