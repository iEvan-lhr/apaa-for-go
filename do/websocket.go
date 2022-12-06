package do

import (
	"github.com/gorilla/websocket"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"log"
	"sync"
)

//var home sync.Map

type Websocket struct {
	home sync.Map
}

//func init() {
//	home = sync.Map{}
//}

func (w *Websocket) AddWebSocketListener(mission chan *anything.Mission, data []any) {
	if &w.home == nil {
		w.home = sync.Map{}
	}
	if _, ok := w.home.Load(data[0]); !ok {
		w.home.Store(data[0], data[1])
		//mission <- &anything.Mission{Name: "SendMessage", Pursuit: []any{"string", "系统公告:" + data[0].(string) + " 已加入聊天室"}}
		<-anything.DoChanN("SendMessage", []any{"string", "系统公告:" + data[0].(string) + " 已加入聊天室"})
		mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{"Websocket连接成功" + data[0].(string)}}
	} else {
		mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{"Websocket连接失败 请检查" + data[0].(string)}}
	}
}

func (w *Websocket) RemoveWebSocketListener(mission chan *anything.Mission, data []any) {
	if value, ok := w.home.Load(data[0]); ok {
		err := value.(*websocket.Conn).Close()
		if err != nil {
			panic(err)
		}
		w.home.Delete(data[0])
	}
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{"已移除Websocket连接" + data[0].(string)}}
}

func (w *Websocket) SendMessage(mission chan *anything.Mission, data []any) {
	switch data[0].(string) {
	case "string":
		w.home.Range(func(key, value any) bool {
			anything.OnceSchedule("SendWebSocket", []any{value, 1, data[1], key})
			return true
		})

	case "struct":
	case "interface":

	}
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{""}}
}

func (w *Websocket) SendWebSocket(data []any) {
	//log.Println(data[0].(*websocket.Conn))
	err := data[0].(*websocket.Conn).WriteMessage(data[1].(int), []byte(data[2].(string)))
	if err != nil {
		log.Println("ERROR:" + err.Error())
		w.home.Delete(data[3])
	}
	//mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{"已发送信息"}}
	//reflect.DeepEqual()

}
