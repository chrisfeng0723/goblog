package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	UserName string
}

type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

type Post struct {
	User
	Body string
}

func PopulateTemplates() map[string]*template.Template {
	const basePath = "view"
	result := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_bash.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Faild to open template blocks directory:" + err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory:" + err.Error())
	}

	for _, fi := range files {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template'" + fi.Name() + "'")
			}

			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file'" + f.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of'" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}()
	}
	return result
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{UserName: "chris"}
		u2 := User{UserName: "rain"}

		posts := []Post{
			Post{User: u1, Body: "i am chris"},
			Post{User: u2, Body: "I am rain"},
		}
		v := IndexViewModel{Title: "Home Page", User: u1, Posts: posts}
		templates :=PopulateTemplates()
		templates["index.html"].Execute(w,&v)
	})
	http.ListenAndServe(":8888", nil)
}
