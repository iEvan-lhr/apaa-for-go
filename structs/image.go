package structs

type UserImage struct {
	ImageId     int64  `json:"image_id"`
	UserId      int64  `json:"user_id"`
	RoomId      int64  `json:"room_id"`
	Image       string `json:"image"`
	Deleted     int64  `json:"deleted"`
	CreatedTime int64  `json:"created_time"`
}
