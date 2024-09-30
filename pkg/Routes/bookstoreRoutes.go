package Routes

import(
	"github.com/gorilla/mux"
	"github.com/ritisha8/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes=func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.Getbook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.getBooksById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}