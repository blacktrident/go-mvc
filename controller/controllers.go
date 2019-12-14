package controller

import (
	"github.com/blacktrident/go-mvc/model"
	"github.com/blacktrident/go-mvc/store"
	"github.com/blacktrident/go-mvc/url"
	"github.com/gorilla/schema"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
		data := struct {
			Title string
		}{
			Title: "My page",
		}
		t, _ := template.ParseFiles("template/add.html")
		t.Execute(res, data)
	}
	if req.Method == "POST" {
		err := req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("BoxArt")
		saveFile(file, handler, err)
		game := new(model.Game)
		decoder := schema.NewDecoder()
		err = decoder.Decode(game, req.Form)
		game.BoxArt = "./src/images/" + handler.Filename
		log.Print(game)
		err = store.Save(game)
		http.Redirect(res, req, urls.ADD_PATH, http.StatusSeeOther)
	}
}
func ShowController(res http.ResponseWriter, req *http.Request) {
	// Get Data from Form and re Execute the query

}

func ShowAllController(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
		Games []model.Game
	}{
		Title: "My page",
	}
	if req.Method == "GET" {
		data.Games, _ = store.GetAll()
		log.Print(data.Games)
		t, _ := template.ParseFiles("template/showAll.html")
		t.Execute(res, data)
	}
}

func saveFile(file multipart.File, handler *multipart.FileHeader, err error) {
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./src/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
