package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
	"net/url"
)

func (o *OpenAPI) GetCharacters(name string) ([]model.CharacterInfo, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	c := []model.CharacterInfo{}

	subUrl := fmt.Sprintf("/characters/%s/siblings", url.QueryEscape(name))
	resp, err := o.SendRequest(http.MethodGet, subUrl, header, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
