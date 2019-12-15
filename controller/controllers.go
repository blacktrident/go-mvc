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
	"strings"
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
		game.Name = strings.ToUpper(game.Name)
		game.Genre = strings.ToUpper(game.Genre)
		game.Platform = strings.ToUpper(game.Platform)
		game.ReleaseDate = strings.ToUpper(game.ReleaseDate)
		log.Print(game)
		err = store.Save(game)
		http.Redirect(res, req, urls.ADD_PATH, http.StatusSeeOther)
	}
}

func ShowController(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		name := req.Form["Name"]
		log.Print(name)
		data := struct {
			Title string
			Game  model.Game
		}{
			Title: "My page",
		}
		game, _ := store.GetOne(strings.ToUpper(name[0]))
		data.Game = *game
		t, _ := template.ParseFiles("template/showOne.html")
		t.Execute(res, data)
	} else {
		log.Print("HereGet")
		data := struct {
			Title string
		}{
			Title: "My page",
		}
		t, _ := template.ParseFiles("template/ShowOne.html")
		t.Execute(res, data)
	}
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
