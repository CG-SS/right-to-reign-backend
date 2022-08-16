package ebs

import "github.com/google/uuid"

type Civilization uint8

const (
	Bohemians Civilization = iota
	Britons
	Franks
	Byzantines
)

type User struct {
	ID  uuid.UUID    `json:"id"`
	Civ Civilization `json:"civ"`
}

type CreateCity struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  int32     `json:"latitude"`
	Longitude int32     `json:"longitude"`
}

type UpdateCity struct {
	ID                 uuid.UUID `json:"id"`
	Name               *string   `json:"name,omitempty"`
	TownCenterLevel    *uint8    `json:"town_center_level,omitempty"`
	HouseLevel         *uint8    `json:"house_level,omitempty"`
	MillLevel          *uint8    `json:"mill_level,omitempty"`
	FarmLevel          *uint8    `json:"farm_level,omitempty"`
	LumberCampLevel    *uint8    `json:"lumber_camp_level,omitempty"`
	MiningCampLevel    *uint8    `json:"mining_camp_level,omitempty"`
	DockLevel          *uint8    `json:"dock_level,omitempty"`
	BarracksLevel      *uint8    `json:"barracks_level,omitempty"`
	PalisadeWallLevel  *uint8    `json:"palisade_wall_level,omitempty"`
	StoneWallLevel     *uint8    `json:"stone_wall_level,omitempty"`
	OutpostLevel       *uint8    `json:"outpost_level,omitempty"`
	MarketLevel        *uint8    `json:"market_level,omitempty"`
	BlacksmithLevel    *uint8    `json:"blacksmith_level,omitempty"`
	ArcheryRangeLevel  *uint8    `json:"archery_range_level,omitempty"`
	StableLevel        *uint8    `json:"stable_level,omitempty"`
	WatchTowerLevel    *uint8    `json:"watch_tower_level,omitempty"`
	UniversityLevel    *uint8    `json:"university_level,omitempty"`
	SiegeWorkshopLevel *uint8    `json:"siege_workshop_level,omitempty"`
	CastleLevel        *uint8    `json:"castle_level,omitempty"`
}
