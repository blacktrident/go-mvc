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
	return router_impl
}
