package openapi

import (
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
	"net/url"
)

const (
	GetCharacters = "/characters/%s/siblings"
)

func (o *OpenAPI) GetCharacters(name string) ([]model.CharacterInfo, error) {
	var resp []model.CharacterInfo
	subUrl := fmt.Sprintf(GetCharacters, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}
