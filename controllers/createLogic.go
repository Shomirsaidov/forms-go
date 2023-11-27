package controllers

import (
	"net/http"
	"../models"
	"../types"
	"fmt"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	formData := types.FormType{r.FormValue("title"), r.FormValue("text"),r.FormValue("author"),r.FormValue("date")}
	
	if r.FormValue("title") != "" && r.FormValue("text") != "" &&
		r.FormValue("author") != "" && r.FormValue("date") != "" {
			models.InsertPost(formData)
			http.Redirect(w ,r , "/", http.StatusSeeOther)
	} else {
		fmt.Println("Blank fields are not allowed !")
		fmt.Fprintf(w, "Blank fields are not allowed !")
	}

}

