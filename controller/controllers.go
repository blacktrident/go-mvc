package controller

import (
	"html/template"
	"net/http"
)

func HomeController(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
	}{
		Title: "My page",
	}
	t, _ := template.ParseFiles("template/home.html")
	t.Execute(res, data)
}
