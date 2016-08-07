package main

import "fmt"
import "net/http"
import "html/template"

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
}

func indexTemplateHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "Index.html")
}

func indexHTMLTemplateHandler(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("hello template")    //create a new template with some name
	tmplt, _ = tmplt.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation

	p := Student{Id: 1, Name: "Mary"} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func indexHTMLTemplateVariableHandler(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("IndexTemplated.html")       //create a new template with some name
	tmplt, _ = tmplt.ParseFiles("IndexTemplated.html") //parse some content and generate a template, which is an internal representation

	p := Student{Name: "Mary"} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func main() {
	fmt.Println("Hello world")

	//http.HandleFunc("/", handler)
	http.HandleFunc("/e", handler2)

	//http.HandleFunc("/", indexTemplateHandler)

	http.HandleFunc("/", indexHTMLTemplateVariableHandler)

	http.ListenAndServe(":8080", nil)
}
