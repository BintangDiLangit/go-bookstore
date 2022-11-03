package routes

import (
	"github.com/BintangDiLangit/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook)
}

// func ()  {

// }
