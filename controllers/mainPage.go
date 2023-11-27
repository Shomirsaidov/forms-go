package controllers

import (
	"net/http"
	"html/template"
	models "../models"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	var posts []models.Post
	posts = models.SelectPosts()

	tmpl, _ := template.ParseFiles("./templates/main.html","./templates/footer.html","./templates/navigation.html")
	tmpl.ExecuteTemplate(w, "main", posts)
}