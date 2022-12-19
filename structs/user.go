package structs

import tools "github.com/iEvan-lhr/exciting-tool"

type User struct {
	// 用户ID
	Id int `json:"id"`
	// 用户名
	Username string `json:"username" marshal:"check"`
	// 用户密码
	Password string `json:"password"`
	// 用户身份认证
	Identity string `json:"identity"`
	// 二维码认证
	QrCode string `json:"qr_code"`
	// 动态Key MD4 加密
	DenKey string `json:"den_key"`
	// 聊天动态Key
	TalkingKey string `json:"talking_key"`
}

type RoomPower struct {
	Id int `json:"id"`
	// 管理房间
	MasterRoom *tools.String `json:"master_room"`
	// 加入房间
	JoinRoom *tools.String `json:"join_room"`
	// 创建房间
	CreatedRoom *tools.String `json:"created_room"`
}

type DenKeyByRegister struct {
	// 注册唯一键Key
	DenKey string `json:"den_key" marshal:"check"`
	// 规则符合判断
	CheckKey string `json:"check_key"`
}
