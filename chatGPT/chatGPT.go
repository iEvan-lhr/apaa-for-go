package chatGPT

import (
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"net/http"
)

type ChatGPT struct {
}

func (c *ChatGPT) SendMessageToChatGPT(mission chan *anything.Mission, data []any) {
	req := data[1].(*http.Request)
	pursuit := (<-anything.DoChanN("TalkToChatGPT", []any{tools.UnMarshal(req, &structs.Talk{})})).Pursuit
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{pursuit}}
}
