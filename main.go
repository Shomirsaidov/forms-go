package main

import (
	"fmt"
	"net/http"
	controllers "./controllers"
	"github.com/gorilla/mux"
)


func main() {
	fmt.Println("Server is working")
	handleRequests()
	
}

func handleRequests() {
	router := mux.NewRouter() 

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", controllers.MainPage).Methods("GET")
	router.HandleFunc("/about/", controllers.AboutPage).Methods("GET")
	router.HandleFunc("/create/", controllers.CreatePage).Methods("GET")
	router.HandleFunc("/createLogic/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/{title}/", controllers.PostPage).Methods("GET")


	http.Handle("/", router)
	http.ListenAndServe(":7070", nil)
}


















