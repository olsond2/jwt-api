package main


import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	myRouter := mux.NewRouter()


	myRouter.Handle("/", http.FileServer(http.Dir("./views/")))
	//myRouter.Handle("/status", NotImplemented).Methods("GET")
	//myRouter.Handle("/products", NotImplemented).Methods("GET")

	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":7000",handlers.LoggingHandler(os.Stdout, myRouter))


}
