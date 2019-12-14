package router

import (
	"github.com/blacktrident/go-mvc/controller"
	"github.com/blacktrident/go-mvc/url"
	"github.com/gorilla/pat"
	"net/http"
)

func GetRouter() *pat.Router {
	urls := url.GetURLS()
	router_impl := pat.New()
	router_impl.Get(urls.SHOW_PATH, controller.ShowController)
	router_impl.Get(urls.SHOWALL_PATH, controller.ShowAllController)
	router_impl.Get(urls.ADD_PATH, controller.AddController)
	router_impl.Post(urls.ADD_PATH, controller.AddController)
	router_impl.PathPrefix(urls.STATIC_PATH).Handler(
		http.StripPrefix(urls.STATIC_PATH, http.FileServer(http.Dir("src"))))
	router_impl.Get(urls.HOME_PATH, controller.HomeController)

	return router_impl
}
