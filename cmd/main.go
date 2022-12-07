package main

import (
	"fmt"
	"github.com/dybrkr/goArk/openapi"
	"github.com/dybrkr/goArk/openapi/model"
	"github.com/dybrkr/goArk/userinfo"
)

const (
	EndPoint = "https://developer-lostark.game.onstove.com"
)

func main() {
	userinfo := userinfo.GetUserInfo()
	fmt.Printf("user => %v\n", userinfo)
	api := openapi.CreateAPI(EndPoint, userinfo.AccessKey)

	resp, err := api.SearchItem(model.GRADE, 0, "", 0, "", "", 0, model.ASC)
	if err != nil {
		return
	}
	fmt.Printf("resp => %v\n", resp)
}
