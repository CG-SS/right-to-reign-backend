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
	Name               *string   `json:"name"`
	TownCenterLevel    *uint8    `json:"town_center_level"`
	HouseLevel         *uint8    `json:"house_level"`
	MillLevel          *uint8    `json:"mill_level"`
	FarmLevel          *uint8    `json:"farm_level"`
	LumberCampLevel    *uint8    `json:"lumber_camp_level"`
	MiningCampLevel    *uint8    `json:"mining_camp_level"`
	DockLevel          *uint8    `json:"dock_level"`
	BarracksLevel      *uint8    `json:"barracks_level"`
	PalisadeWallLevel  *uint8    `json:"palisade_wall_level"`
	StoneWallLevel     *uint8    `json:"stone_wall_level"`
	OutpostLevel       *uint8    `json:"outpost_level"`
	MarketLevel        *uint8    `json:"market_level"`
	BlacksmithLevel    *uint8    `json:"blacksmith_level"`
	ArcheryRangeLevel  *uint8    `json:"archery_range_level"`
	StableLevel        *uint8    `json:"stable_level"`
	WatchTowerLevel    *uint8    `json:"watch_tower_level"`
	UniversityLevel    *uint8    `json:"university_level"`
	SiegeWorkshopLevel *uint8    `json:"siege_workshop_level"`
	CastleLevel        *uint8    `json:"castle_level"`
}
