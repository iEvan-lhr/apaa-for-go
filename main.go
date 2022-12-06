package main

import (
	"cloud-deer/do"
	"cloud-deer/model"
	"cloud-deer/router"
	"cloud-deer/structs"
	"cloud-deer/websocket"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"github.com/iEvan-lhr/worker/engine"
)

func main() {
	e := engine.Engine{
		W:      anything.Wind{},
		Origin: []string{"*", "POST, GET, OPTIONS, PUT, DELETE", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
	}
	e.Start("9080", []any{&do.Websocket{}, &structs.DBConn{}, &model.UserModel{}}, []any{&router.Router{}, &websocket.Dad{}},
		//初始化需执行的方法
		map[string][]any{
			"GetConn": {"mysql", "root:Luhaoran0!@tcp(106.12.170.224:3306)/evan?parseTime=true"},
		})
}
