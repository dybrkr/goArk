package openapi

import (
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

func (o *OpenAPI) GetProfiles(name string) (model.ArmoryProfile, error) {
	var resp model.ArmoryProfile

	subUrl := fmt.Sprintf(GetProfilesFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetEquipment(name string) ([]model.ArmoryEquipment, error) {
	var resp []model.ArmoryEquipment
	subUrl := fmt.Sprintf(GetEquipmentFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetAvatars(name string) (model.ArmoryAvatars, error) {
	var resp model.ArmoryAvatars
	subUrl := fmt.Sprintf(GetAvatarFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetCombatSkills(name string) ([]model.ArmorySkill, error) {
	var resp []model.ArmorySkill
	subUrl := fmt.Sprintf(GetCombatSkillFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetEngravings(name string) (model.ArmoryEngraving, error) {
	var resp model.ArmoryEngraving
	subUrl := fmt.Sprintf(GetEngravingsFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetCards(name string) (model.ArmoryCard, error) {
	var resp model.ArmoryCard
	subUrl := fmt.Sprintf(GetCardFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetGems(name string) (model.ArmoryGem, error) {
	var resp model.ArmoryGem
	subUrl := fmt.Sprintf(GetGemsFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetColosseums(name string) (model.ColosseumInfo, error) {
	var resp model.ColosseumInfo
	subUrl := fmt.Sprintf(GetColosseumsFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetCollectibles(name string) ([]model.Collectible, error) {
	var resp []model.Collectible
	subUrl := fmt.Sprintf(GetCollectiblesFMT, url.QueryEscape(name))
	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}
