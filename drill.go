// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"html/template"
// )

// type User struct {
// 	Name string
// 	Age uint16
// 	Score float64
// 	Friends []string
// }

// func(u User) getInfo() string {
// 	return fmt.Sprintf("<b>Name : %s. <br/>Age : %d. <br/>Score : %f</b>", u.Name, u.Age, u.Score)
// }

// func(u *User) setName(name string) {
// 	u.Name = name
// }





// func mainPage(w http.ResponseWriter, r *http.Request) {
// 	abubakr := User{"Abubakr",15,0.99,[]string{"Bilol","Akmal","Firdavs"}}
// 	abubakr.setName("Shomirsaidov")

// 	// fmt.Fprintf(w, abubakr.getInfo())
// 	tmpl, _ := template.ParseFiles("templates/home.html")
// 	tmpl.Execute(w, abubakr)


// } 

// func aboutPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "About page")
// } 

// func handleRequests() {
// 	http.HandleFunc("/", mainPage)
// 	http.HandleFunc("/about/", aboutPage)

// 	http.ListenAndServe(":7070", nil)
// }

// func main() {
// 	fmt.Println("Server is working")
// 	handleRequests()
	
// }

