package main

import (
	"github.com/blacktrident/go-mvc/router"
	"github.com/urfave/negroni"
	"log"
)

func main() {
	router := router.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening:")
	n.Run(":3001")
	//var x model.Game
	//x.Name = "test"
	//x.Genre = "test"
	//store.SaveGame(&x)
	//var x []string
	//x = append(x, "efwf")
	//z, _ := store.GetOne("efwf")
	//log.Print(z)
}
