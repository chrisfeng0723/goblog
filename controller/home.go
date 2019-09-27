package controller

import (
	"github.com/chrisfeng0723/goblog/vm"
	"net/http"
)

type home struct {
}

func (h home) registerRouters() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVm()
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 {
			v.AddError("username must longer than s")
		}

		if len(password) < 4 {
			v.AddError("password must longer than 4")
		}

		if !check(username,password){
			v.AddError("username password not correct,pls input again")
		}
		if len(v.Errs) >0{
			templates[tpName].Execute(w,&v)
		}else{
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}
	}

}

func check(username, password string) bool {
	if username == "chris" && password == "rain" {
		return true
	}
	return false
}
