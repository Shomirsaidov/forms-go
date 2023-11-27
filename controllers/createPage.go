package controllers

import (
	"net/http"
	"html/template"
)

func CreatePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/create.html")
	tmpl.ExecuteTemplate(w, "create","Success !")


}