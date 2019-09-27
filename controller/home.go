package controller

import (
	"github.com/chrisfeng0723/goblog/vm"
	"net/http"
)

type home struct {

}

func(h home) registerRouters(){
	http.HandleFunc("/",indexHandler)
}

func indexHandler(w http.ResponseWriter,r *http.Request){
	vop :=vm.IndexViewModelOp{}
	v:=vop.GetVM()
   templates["index.html"].Execute(w,&v)
}
