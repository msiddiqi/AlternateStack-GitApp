package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "Index.html")
}

func TemplatedHandler(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("hello template")
	tmplt, _ = tmplt.Parse("Top Student: {{.Id}} - {{.Name}}!")

	p := Student{Id: 1, Name: "Aisha"} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func main() {

	fmt.Println("Server Starting")
	http.HandleFunc("/", SimpleIndexHandler)
	http.HandleFunc("/index", HttpFileHandler)

	http.HandleFunc("/top-student", TemplatedHandler)
	//http.HandleFunc("/", indexTemplateHandler)

	http.ListenAndServe(":8080", nil)
}
