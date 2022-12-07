/*
Key 정보를 따로 저장하기 위한 임시 패키지

같은 디렉토리에 userinfo.json 에 저장된 Key를 가져와서 사용
*/
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
