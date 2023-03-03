package model

import (
	"log"
	"testing"
)

func TestChatGPT(t *testing.T) {
	log.Println(askChatGPT("system", "你好"))
}
