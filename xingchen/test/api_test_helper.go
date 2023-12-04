package xingchen

import (
	"fmt"
	"io"
	"net/http"
)

func loadV1Token() string {
	// user id ce8f9557131d499a869c777eafddd879
	resp, err := http.Get("https://alimt-plugins.alibaba.com/api/auth/refresh?token=XXX")
	if err != nil {
		fmt.Println("Request failed:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body fail:", err)
		return ""
	}
	return string(body)
}
