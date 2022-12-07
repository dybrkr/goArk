package model

import "time"

type GuildRanking struct {
	Rank           int       `json:"Rank"`
	GuildName      string    `json:"GuildName"`
	GuildMessage   string    `json:"GuildMessage"`
	MasterName     string    `json:"MasterName"`
	Rating         int       `json:"Rating"`
	MemberCount    int       `json:"MemberCount"`
	MaxMemberCount int       `json:"MaxMemberCount"`
	UpdatedDate    time.Time `json:"UpdatedDate"`
}
