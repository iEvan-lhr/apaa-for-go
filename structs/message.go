package structs

type Message struct {
	// 时间戳ID
	// 房间ID
	// 信息详情 []HotMessage
	// 物理删除
	// 创建时间
	// 修改时间
}

type HotMessage struct {
	// 用户
	// 内容
	// 时间戳ID
	// 房间ID
	// 信息可见性 []User 为空默认房间全部可见
}
