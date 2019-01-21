package main
import (
	"html/template"
	"os"
)

func main() {
	tmpl = template
	tmpl, err := template.Must("./hello.template")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, "World" )
	if err != nil { panic(err) }
}
