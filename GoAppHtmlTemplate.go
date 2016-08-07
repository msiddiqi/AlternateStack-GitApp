package main

import "fmt"
import "net/http"
import "html/template"

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

func indexHTMLTemplateVariableHandler(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("IndexTemplated.html")       //create a new template with some name
	tmplt, _ = tmplt.ParseFiles("IndexTemplated.html") //parse some content and generate a template, which is an internal representation

	p := Student{Id: 1, Name: "Aisha"} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func main() {
	fmt.Println("Starting Server for Templated response from file")

	http.HandleFunc("/top-student-from-file", indexHTMLTemplateVariableHandler)

	http.ListenAndServe(":8080", nil)
}
