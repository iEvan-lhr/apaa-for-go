package model

import (
	"crypto/tls"
	"github.com/iEvan-lhr/apaa-for-go/structs"
	build "github.com/iEvan-lhr/encryption-apaa"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ChatGPTModel struct {
}

func (c *ChatGPTModel) TalkToChatGPT(mission chan *anything.Mission, data []any) {
	talk := data[0].(*structs.Talk)
	mission <- &anything.Mission{Pursuit: []any{string(tools.Marshal(structs.UserRes{Status: 200, Ans: askChatGPT("system", talk.Say)}))}}
}

func askChatGPT(role, message string) string {
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", build.ReturnAuthKey())
	header.Set("Accept", "*/*")
	header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Type", "application/json")
	header.Set("User-Agent", "PostmanRuntime/7.28.4")
	mes := `{"model": "gpt-3.5-turbo-0301","messages": [{"role": "` + role + `", "content": "` + message + `"}]}`
	res := &structs.ChatGPTRes{}
	tr := &http.Transport{
		Proxy: http.ProxyURL(tools.ReturnValue(url.Parse("http://127.0.0.1:10001/")).(*url.URL)),
		// 忽略不受信任的 TLS 证书
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req := tools.ReturnValue(http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(mes))).(*http.Request)
	req.Header = header
	tools.Unmarshal(tools.DoReq(req, &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}).String(), res)
	return res.Choices[0].Message.Content
}
