package controller

import (
	"github.com/blacktrident/go-mvc/model"
	"github.com/blacktrident/go-mvc/store"
	"github.com/blacktrident/go-mvc/url"
	"github.com/gorilla/schema"
	"html/template"
	"log"
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

func AddController(res http.ResponseWriter, req *http.Request) {
	//data := make(map[string]interface{})
	//controllerTemplate := templates.POST_ADD
	urls := url.GetURLS()
	if req.Method == "GET" {
		//t, _ := template.ParseFiles("template/add.html")
		//res.
		//return add.html
	}
	if req.Method == "POST" {
		err := req.ParseForm()
		game := new(model.Game)
		decoder := schema.NewDecoder()
		err = decoder.Decode(game, req.Form)
		log.Println(err)
		if err != nil {
			//utils.CustomTemplateExecute(res, req, controllerTemplate, data)
		} else {
			err = store.SaveGame(game)
			http.Redirect(res, req, urls.ADD_PATH, http.StatusSeeOther)
		}
	}
}
