package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"net/http"
)

type Dad struct {
}

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
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
	anything.OnceSchedule("ReadMessage", []any{ws, data[1].(*http.Request).FormValue("name")})
	mission <- &anything.Mission{Name: "-AddWebSocketListener", Pursuit: []any{data[1].(*http.Request).FormValue("name"), ws}}
}

func (d *Dad) SendWebsocketMessage(mission chan *anything.Mission, data []any) {
	req := data[1].(*http.Request)
	talk := tools.UnMarshal(req, &structs.Talk{}).(*structs.Talk)
	switch talk.TalkType {
	case "T":
		mission <- &anything.Mission{Name: "-SendMessage", Pursuit: []any{"string", string(tools.Marshal(talk))}}
	case "I":
		image := &structs.UserImage{
			ImageID: int64(tools.ReturnValue(tools.Make(talk.Say).Atoi()).(int)),
		}
		structs.Find(image, &image)
		talk.Say = image.Image
		mission <- &anything.Mission{Name: "-SendMessage", Pursuit: []any{"string", string(tools.Marshal(talk))}}
	}

}

func (d *Dad) JoinRoom(mission chan *anything.Mission, data []any) {

}
