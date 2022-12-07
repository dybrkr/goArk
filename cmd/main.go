package main

import (
	"fmt"
	"github.com/dybrkr/goArk/request"
	"github.com/dybrkr/goArk/userinfo"
	"net/http"
)

const (
	EndPoint = "https://developer-lostark.game.onstove.com"
)

func main() {
	userinfo := userinfo.GetUserInfo()
	fmt.Printf("user => %v\n", userinfo)
	api := request.CreateAPI(EndPoint)

	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + userinfo.AccessKey

	resp, err := api.SendAuthRequest(http.MethodGet, "/news/events", header, nil)
	if err != nil {
		return
	}
	fmt.Printf("resp => %v\n", resp)
}
