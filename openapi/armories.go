package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
	"net/url"
)

const (
	GetProfilesFMT     = "/armories/characters/%s/profiles"
	GetEquipmentFMT    = "/armories/characters/%s/equipment"
	GetAvatarFMT       = "/armories/characters/%s/avatars"
	GetCombatSkillFMT  = "/armories/characters/%s/combat-skills"
	GetEngravingsFMT   = "/armories/characters/%s/engravings"
	GetCardFMT         = "/armories/characters/%s/cards"
	GetGemsFMT         = "/armories/characters/%s/gems"
	GetColosseumsFMT   = "/armories/characters/%s/colosseums"
	GetCollectiblesFMT = "/armories/characters/%s/collectibles"
)

type ArmoriesResponse interface {
	model.ArmoryProfile | []model.ArmoryEquipment | model.ArmoryAvatars |
		[]model.ArmorySkill | model.ArmoryEngraving | model.ArmoryCard |
		model.ArmoryGem | model.ColosseumInfo | []model.Collectible
}

func SendArmoriesRequest[T ArmoriesResponse](o *OpenAPI, urlPath string, name string, out *T) error {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	subUrl := fmt.Sprintf(urlPath, url.QueryEscape(name))
	resp, err := o.SendRequest(http.MethodGet, subUrl, header, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(resp), out)
	if err != nil {
		return err
	}

	return nil
}

func (o *OpenAPI) GetProfiles(name string) (model.ArmoryProfile, error) {
	var resp model.ArmoryProfile
	err := SendArmoriesRequest(o, GetProfilesFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetEquipment(name string) ([]model.ArmoryEquipment, error) {
	var resp []model.ArmoryEquipment
	err := SendArmoriesRequest(o, GetEquipmentFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetAvatars(name string) (model.ArmoryAvatars, error) {
	var resp model.ArmoryAvatars
	err := SendArmoriesRequest(o, GetAvatarFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetCombatSkills(name string) ([]model.ArmorySkill, error) {
	var resp []model.ArmorySkill
	err := SendArmoriesRequest(o, GetCombatSkillFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetEngravings(name string) (model.ArmoryEngraving, error) {
	var resp model.ArmoryEngraving
	err := SendArmoriesRequest(o, GetEngravingsFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetCards(name string) (model.ArmoryCard, error) {
	var resp model.ArmoryCard
	err := SendArmoriesRequest(o, GetCardFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetGems(name string) (model.ArmoryGem, error) {
	var resp model.ArmoryGem
	err := SendArmoriesRequest(o, GetGemsFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetColosseums(name string) (model.ColosseumInfo, error) {
	var resp model.ColosseumInfo
	err := SendArmoriesRequest(o, GetColosseumsFMT, name, &resp)
	return resp, err
}

func (o *OpenAPI) GetCollectibles(name string) ([]model.Collectible, error) {
	var resp []model.Collectible
	err := SendArmoriesRequest(o, GetCollectiblesFMT, name, &resp)
	return resp, err
}
