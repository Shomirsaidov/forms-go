package controllers

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tmpl, _ := template.ParseFiles("./templates/post.html")
	tmpl.ExecuteTemplate(w, "post",vars["title"])


}

