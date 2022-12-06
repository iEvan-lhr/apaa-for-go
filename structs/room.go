package structs

import "time"

type Room struct {
	// 房间ID
	RoomID int64 `json:"room_id"`
	// 房间名称
	RoomName string `json:"room_name"`
	// 房间状态
	RoomStatus string `json:"room_status"`
	// 房间用户数量
	RoomUserList string `json:"room_user_list"`
	// 是否可被检索
	InIndex int `json:"in_index"`
	// 内容保存周期
	SaveTime int64 `json:"save_time"`
	// 是否私链
	Private bool `json:"private"`
	// 是否允许子房间
	InnerRoom bool `json:"inner_room"`
	// 物理删除
	Deleted bool `json:"deleted"`
	// 创建时间
	CreatedTime time.Time `json:"created_time"`
	// 修改时间
	UpdatedTime time.Time `json:"updated_time"`
}
