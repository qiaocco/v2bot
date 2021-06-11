package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestDiff(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	res := difference(a, b)
	fmt.Println(res)
}

func TestSendMsg(t *testing.T) {
	token := os.Getenv("HedwigToken")

	reqBody := &sendMessageReqBody{
		ChatID: "@qiaocc_pushbot",
		Text:   "你好",
	}
	reqBytes, _ := json.Marshal(reqBody)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err)
	}

}
