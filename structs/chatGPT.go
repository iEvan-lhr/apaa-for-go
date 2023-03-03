package structs

type ChatGPTRes struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   Usage     `json:"usage"`
	Choices []Choices `json:"choices"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choices struct {
	Message      MessageChatGPT `json:"message"`
	FinishReason string         `json:"finish_reason"`
	Index        int            `json:"index"`
}
type MessageChatGPT struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
