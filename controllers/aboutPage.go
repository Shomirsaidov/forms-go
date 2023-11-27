package controllers

import (
	"net/http"
	"html/template"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	var Info = Data{[]string{"Abubakr","Bilol"}}
	tmpl, _ := template.ParseFiles("./templates/about.html","./templates/footer.html","./templates/navigation.html")
	tmpl.ExecuteTemplate(w, "about", Info)
}

type Data struct {
	Names []string
}