package main

import (
	"github.com/chrisfeng0723/goblog/controller"
	"net/http"
)



func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
