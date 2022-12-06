package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"net/http"
)

type Dad struct {
}

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (d *Dad) GetDadConn(mission chan *anything.Mission, data []any) {
	ws, err := upGrader.Upgrade(data[0].(http.ResponseWriter), data[1].(*http.Request), nil)
	if err != nil {
		panic(err)
	}
	mission <- &anything.Mission{Name: "-AddWebSocketListener", Pursuit: []any{data[1].(*http.Request).FormValue("name"), ws}}
}

func (d *Dad) SendWebsocketMessage(mission chan *anything.Mission, data []any) {
	mission <- &anything.Mission{Name: "SendMessage", Pursuit: []any{"string", data[1].(*http.Request).FormValue("message")}}
}

func (d *Dad) JoinRoom(mission chan *anything.Mission, data []any) {

}
