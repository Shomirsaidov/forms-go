package models

// import (
// 	"log"
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"
// )



// var posts = []Post{}

// func SelectPosts() []Post {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
//     defer db.Close()

//     if err != nil {
//         log.Fatal(err)
//     }

//     data, err2 := db.Query("SELECT * FROM `posts` ORDER BY `id` DESC")

//     posts = []Post{}
//     for data.Next() {
//         var post Post
//         err = data.Scan(&post.Id, &post.Title, &post.Text,&post.Author, &post.Date)

//         if err != nil {
//             panic(err)
//         }

//         posts = append(posts, post)
//     }

//     if err2 != nil {
//         log.Fatal(err2)
//     }

    
//     return posts
// }













