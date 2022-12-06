package structs

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Identity   string `json:"identity"`
	QrCode     string `json:"qr_code"`
	DenKey     string `json:"den_key"`
	TalkingKey string `json:"talking_key"`
}
