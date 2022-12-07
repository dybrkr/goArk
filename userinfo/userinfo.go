package userinfo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type UserInfo struct {
	UserName  string `json:"userName"`
	AccessKey string `json:"AccessKey"`
}

const (
	defaultConfigPath = "userinfo.json"
)

var (
	u = load()
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func load() UserInfo {
	ui := UserInfo{}
	isExist := fileExists(defaultConfigPath)
	if isExist == true {
		data, err := ioutil.ReadFile(defaultConfigPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &ui)
		if err != nil {
			panic(err)
		}
	} else {
		jsonBytes, _ := json.Marshal(ui)
		ioutil.WriteFile(defaultConfigPath, jsonBytes, 0644)
	}
	return ui
}

func GetUserInfo() UserInfo {
	return u
}
