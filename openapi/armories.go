package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
	"net/url"
)

const (
	GetProfilesFMT  = "/armories/characters/%s/profiles"
	GetEquipmentFMT = "/armories/characters/%s/equipment"
	GetAvatarFMT    = "/armories/characters/%s/avatars"
)

func (o *OpenAPI) GetProfiles(name string) (model.ArmoryProfile, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	p := model.ArmoryProfile{}

	subUrl := fmt.Sprintf(GetProfilesFMT, url.QueryEscape(name))
	resp, err := o.SendAuthRequest(http.MethodGet, subUrl, header, nil)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (o *OpenAPI) GetEquipment(name string) (model.ArmoryEquipment, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	p := model.ArmoryEquipment{}

	subUrl := fmt.Sprintf(GetEquipmentFMT, url.QueryEscape(name))
	resp, err := o.SendAuthRequest(http.MethodGet, subUrl, header, nil)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (o *OpenAPI) GetAvatars(name string) (model.ArmoryAvatar, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	p := model.ArmoryAvatar{}

	subUrl := fmt.Sprintf(GetAvatarFMT, url.QueryEscape(name))
	resp, err := o.SendAuthRequest(http.MethodGet, subUrl, header, nil)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
