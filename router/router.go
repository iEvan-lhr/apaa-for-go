package router

import (
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"net/http"
	"xorm.io/xorm"
)

type Router struct {
}

func (r *Router) TestRouter(mission chan *anything.Mission, data []any) {
	engine := (<-anything.DoChanN("GetConn", []any{"mysql", "root:Luhaoran0!@tcp(106.12.170.224:3306)/evan?parseTime=true"})).Pursuit[0].(*xorm.Engine)
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{engine.Ping()}}
}

func (r *Router) UserLogin(mission chan *anything.Mission, data []any) {
	req := data[1].(*http.Request)
	pursuit := (<-anything.DoChanN("Login", []any{tools.UnMarshal(req, &structs.User{})})).Pursuit
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{pursuit}}
}

func (r *Router) UserRegister(mission chan *anything.Mission, data []any) {
	req := data[1].(*http.Request)
	pursuit := (<-anything.DoChanN("Register", []any{tools.UnMarshal(req, &structs.User{})})).Pursuit
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{pursuit}}
}

func (r *Router) FindImageByUserAndRoom(mission chan *anything.Mission, data []any) {
	req := data[1].(*http.Request)
	pursuit := (<-anything.DoChanN("GetImages", []any{tools.UnMarshal(req, &structs.UserImage{})})).Pursuit
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{pursuit}}
}
