package structs

type UserImage struct {
	ImageID     int64  `json:"image_id"`
	UserID      int64  `json:"user_id"`
	RoomID      int64  `json:"room_id"`
	Image       string `json:"image"`
	Deleted     int64  `json:"deleted"`
	CreatedTime int64  `json:"created_time"`
}
