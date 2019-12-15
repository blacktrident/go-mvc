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
	n.Run(":3000")
}
