package items

// ItemInfo is a struct that contains the general information of an item

type Weapon struct {
	Damage               int64          `json:"damage"`
	CriticalStrikeChance int64          `json:"critical_strike_chance"`
	Knockback            int64          `json:"knockback"`
	ManaCost             int64          `json:"mana_cost"`
	Velocity             int64          `json:"velocity"`
	Minion               int64          `json:"minion"`
	ToolTip              string         `json:"tool_tip"`
	GeneralInfo          []*GeneralInfo `json:"general_info"`
}

type Tool struct {
	GeneralInfo          GeneralInfo `json:"general_info"`
	CriticalStrikeChance int64       `json:"critical_strike_chance"`
	ToolSpeed            int64       `json:"tool_speed"`
	Damage               int64       `json:"damage"`
	Power                int64       `json:"power"`
	MinableOres          []string    `json:"minable_ores"`
}

// Armor items
type Armor struct {
	ArmorType string `json:"armor_type"`
	Defense   int64  `json:"defense"`
	Effect    string `json:"effect"`
	Class     string `json:"class"`
}

type accessory struct {
}

type pets struct {
}
type mount struct {
}
