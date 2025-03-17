// Package stalcraft provides an API client for the [STALCRAFT: X API].
//
// [STALCRAFT: X API]: https://eapi.stalcraft.net/overview.html

package stalcraft

import (
	"time"
)

const (
	RegionNA  = "NA"
	RegionEU  = "EU"
	RegionSEA = "SEA"
	RegionRU  = "RU"
)

type Lot struct {
	ItemID       string         `json:"itemId"`
	Amount       int            `json:"amount"`
	StartPrice   int            `json:"startPrice"`
	CurrentPrice *int           `json:"currentPrice"`
	BuyoutPrice  *int           `json:"buyoutPrice"`
	StartTime    time.Time      `json:"startTime"`
	EndTime      time.Time      `json:"endTime"`
	Additional   map[string]any `json:"additional"`
}

type LotListing struct {
	Total int
	Lots  []Lot
}

type PriceEntry struct {
	Amount     int            `json:"amount"`
	Price      int            `json:"price"`
	Time       time.Time      `json:"time"`
	Additional map[string]any `json:"additional"`
}

type PricesListing struct {
	Total  int          `json:"total"`
	Prices []PriceEntry `json:"prices"`
}

type CharacterClanInfo struct {
	Info   ClanInfo   `json:"info"`
	Member ClanMember `json:"member"`
}
type CharacterMetaInfo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CreationTime time.Time `json:"creationTime"`
}
type CharacterProfileData struct {
	Username              string               `json:"username"`
	UUID                  string               `json:"uuid"`
	Status                string               `json:"status"`
	Alliance              string               `json:"alliance"`
	LastLogin             time.Time            `json:"lastLogin"`
	DisplayedAchievements []string             `json:"displayedAchievements"`
	Clan                  CharacterClanInfo    `json:"clan"`
	Stats                 []CharacterStatValue `json:"stats"`
}

type CharacterStatValue struct {
	ID    string         `json:"id"`
	Type  StatType       `json:"type"`
	Value map[string]any `json:"value"`
}

type FullCharacterInfo struct {
	Information CharacterMetaInfo `json:"information"`
	Clan        CharacterClanInfo `json:"clan"`
}

type StatType string

const (
	StatTypeInteger  StatType = "INTEGER"
	StatTypeDecimal  StatType = "DECIMAL"
	StatTypeDate     StatType = "DATE"
	StatTypeDuration StatType = "DURATION"
)

type ClansListResponse struct {
	TotalClans int        `json:"totalClans"`
	Data       []ClanInfo `json:"data"`
}

type ClanInfo struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Level            int       `json:"level"`
	RegistrationTime time.Time `json:"registrationTime"`
	Alliance         string    `json:"alliance"`
	Description      string    `json:"description"`
	Leader           string    `json:"leader"`
	MemberCount      int       `json:"memberCount"`
}

type ClanMember struct {
	Name     string   `json:"name"`
	Rank     ClanRank `json:"rank"`
	JoinTime string   `json:"joinTime"`
}

type ClanRank string

const (
	ClanRankRecruit  ClanRank = "RECRUIT"
	ClanRankCommoner ClanRank = "COMMONER"
	ClanRankSoldier  ClanRank = "SOLDIER"
	ClanRankSergeant ClanRank = "SERGEANT"
	ClanRankOfficer  ClanRank = "OFFICER"
	ClanRankColonel  ClanRank = "COLONEL"
	ClanRankLeader   ClanRank = "LEADER"
)

type EmissionResponse struct {
	CurrentStart  time.Time `json:"currentStart"`
	PreviousStart time.Time `json:"previousStart"`
	PreviousEnd   time.Time `json:"previousEnd"`
}

type RegionInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
