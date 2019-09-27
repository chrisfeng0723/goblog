package vm

import "github.com/chrisfeng0723/goblog/model"

type IndexViewModel struct {
	BaseViewModel
	model.User

	Posts []model.Post
}

type IndexViewModelOp struct {}

func(IndexViewModelOp) GetVM() IndexViewModel{
	u1 :=model.User{UserName:"chris"}
	u2 :=model.User{UserName:"rain"}

	posts := []model.Post{
		{User:u1,Body:"i am king"},
		{User:u2,Body:"beautiful girl"},
	}

	v := IndexViewModel{BaseViewModel{Title:"HomePage"},u1,posts}

	return v
}
