package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

// render page of html
func RenderPage(page string, res http.ResponseWriter) {
	// Create a new template and add a custom function for comparison
	funcMap := template.FuncMap{
		"eq": func(a, b int) bool {
			return a == b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"lt": func(a, b int) bool {
			return a < b
		},
	}
	temp, err := template.ParseFiles("templates/" + page + ".html")
	temp.Funcs(funcMap)
	if err != nil {
		fmt.Println(err)
		if page == "error" {
			http.Error(res, "Internal Server Error", 500)
			return
		}
		Error(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err1 := temp.Execute(res, Data)
	if err1 != nil {
		fmt.Println(err1.Error())
		if page == "error" {
			http.Error(res, "Internal Server Error", 500)
			return
		}
		Error(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
