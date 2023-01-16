package structs

type Response struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type UserRes struct {
	Status   int    `json:"status"`
	Ans      string `json:"ans"`
	Identity string `json:"identity"`
	Id       int    `json:"id"`
}

type ImageRes struct {
	Status int    `json:"status"`
	Ans    string `json:"ans"`
	Data   string `json:"data"`
}
