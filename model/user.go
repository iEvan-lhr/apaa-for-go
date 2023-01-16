package model

import (
	"github.com/iEvan-lhr/apaa-for-go/structs"
	build "github.com/iEvan-lhr/encryption-apaa"
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
		user[0].DenKey, user[0].Identity = build.Encryption(time.Now().Add(15 * time.Minute).Unix())
		mission <- &anything.Mission{Pursuit: []any{string(tools.Marshal(structs.UserRes{Status: 200, Ans: "Succ", Identity: user[0].Identity, Id: user[0].Id}))}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{string(tools.Marshal(structs.UserRes{Status: 202, Ans: "Fail"}))}}
	}
}

func (u *UserModel) Check(mission chan *anything.Mission, user *structs.User) {
	if ok, _ := build.Decrypt(tools.Strings(user.DenKey), tools.Strings(user.Identity)); !ok {
		mission <- &anything.Mission{Pursuit: []any{false}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{true}}
	}
}

func (u *UserModel) Register(mission chan *anything.Mission, data []any) {
	user := data[0].(*structs.User)
	if u.checkUserLegitimate(user) {
		structs.Save(user)
		mission <- &anything.Mission{Pursuit: []any{"RegisterSuccess"}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{"RegisterFail"}}
	}
}
