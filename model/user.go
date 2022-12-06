package model

import (
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"time"
)

type UserModel struct {
}

func (u *UserModel) Login(mission chan *anything.Mission, data []any) {
	var user []structs.User
	login := data[0].(*structs.User)
	structs.Find(login, &user)
	if len(user) == 1 {
		user[0].DenKey, user[0].Identity = encryption(time.Now().Add(15 * time.Minute).Unix())
		mission <- &anything.Mission{Pursuit: []any{"loginSucc", user[0].Identity}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{"loginFail"}}
	}
}

func (u *UserModel) Check(mission chan *anything.Mission, user *structs.User) {
	if ok, _ := decrypt(tools.Strings(user.DenKey), tools.Strings(user.Identity)); !ok {
		mission <- &anything.Mission{Pursuit: []any{false}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{true}}
	}
}
