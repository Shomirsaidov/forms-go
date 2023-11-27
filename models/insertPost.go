package models

import (
	"fmt"
	"../types"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func InsertPost(formData types.FormType) {
	fmt.Println(formData)

	db, conErr := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	if conErr != nil {
		panic("Connection has failed (")
	}

	_, queryErr := db.Query(fmt.Sprintf("INSERT INTO `posts`(`id`,`title`,`text`,`author`,`date`) VALUES(NULL, '%s', '%s', '%s', '%s')", 
		formData.Title, formData.Text, formData.Author, formData.Date))

	if queryErr != nil {
		panic("Connection has failed (")
	} else {
		fmt.Println("New post !!!")
	}
}










