package structs

type Response struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type UserRes struct {
	Status   int    `json:"status"`
	Ans      string `json:"ans"`
	Identity string `json:"identity"`
}
