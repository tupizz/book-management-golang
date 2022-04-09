package routes

import (
	"book-management-system-golang/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)
}
