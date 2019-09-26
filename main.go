package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	UserName string
}

type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

type Post struct {
	User User
	Body string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{UserName: "chris"}
		u2 :=User{UserName:"rain"}

		posts :=[]Post{
			Post{User:u1,Body:"i am chris"},
			Post{User:u2,Body:"I am rain"},
		}
		v := IndexViewModel{Title: "Home Page", User: u1,Posts:posts}
		tpl, _ := template.ParseFiles("view/index.html")
		fmt.Println(v)
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
