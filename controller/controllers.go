package controller

import (
	"github.com/blacktrident/go-mvc/model"
	"github.com/gorilla/schema"
	"html/template"
	"log"
	"net/http"
	"time"
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

func AddPostController(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	controllerTemplate := templates.POST_ADD
	url_patterns := urls.ReturnURLS()
	if req.Method == "GET" {
		//utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
	if req.Method == "POST" {
		err := req.ParseForm()
		game := new(model.Game)
		decoder := schema.NewDecoder()
		err = decoder.Decode(game, req.Form)
		log.Println(err)
		if err != nil {
			utils.CustomTemplateExecute(res, req, controllerTemplate, data)
		} else {
			http.Redirect(res, req, url_patterns.POSTS_PATH, http.StatusSeeOther)
		}
	}
}
