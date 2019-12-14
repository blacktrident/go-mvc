package router

import (
	"github.com/blacktrident/go-mvc/controller"
	"github.com/blacktrident/go-mvc/url"
	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	urls := url.GetURLS()
	router_impl := pat.New()
	router_impl.Get(urls.Home_PATH, controller.HomeController)
	router_impl.Get(urls.ADD_PATH, controller.AddController)
	router_impl.Post(urls.ADD_PATH, controller.AddController)
	return router_impl
}
