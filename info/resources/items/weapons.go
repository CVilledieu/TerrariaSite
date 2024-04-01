package items

// Equipment items
// Equipment items are items that can be equipped by the player character. Also includes Weapons

type weapon struct {
	Type                 string `json:"type"`
	Damage               int64  `json:"damage"`
	CriticalStrikeChance int64  `json:"critical_strike_chance"`
	Knockback            int64  `json:"knockback"`
	ManaCost             int64  `json:"mana_cost"`
	Velocity             int64  `json:"velocity"`
	Minion               int64  `json:"minion"`
	ToolTip              string `json:"tool_tip"`
}

type tool struct {
	ToolSpeed   int64    `json:"tool_speed"`
	ToolDamage  int64    `json:"tool_damage"`
	Power       int64    `json:"power"`
	MinableOres []string `json:"minable_ores"`
}
