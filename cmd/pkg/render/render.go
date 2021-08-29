package render

import (
	"fmt"
	"net/http"
	"text/template"
)

//This renders the templates that were created, this func is called in the corressponding page functions
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}
