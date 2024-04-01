package items

// Equipment items
// Equipment items are items that can be equipped by the player character. Also includes Weapons
type weapon struct {
	Melee    *meleeWeapon    `json:"melee"`
	Ranged   *rangedWeapon   `json:"ranged"`
	Magic    *magicWeapon    `json:"magic"`
	Summoner *summonerWeapon `json:"summoner"`
	Rouge    *rougeWeapon    `json:"rouge"`
}

type tool struct {
	ToolSpeed   int64    `json:"tool_speed"`
	ToolDamage  int64    `json:"tool_damage"`
	Power       int64    `json:"power"`
	MinableOres []string `json:"minable_ores"`
}

// Weapon items
type meleeWeapon struct {
	Damage               int64 `json:"damage"`
	CriticalStrikeChance int64 `json:"critical_strike_chance"`
	Knockback            int64 `json:"knockback"`
}

type rangedWeapon struct {
}

type magicWeapon struct {
	ManaCost  int64 `json:"mana_cost"`
	Damage    int64 `json:"damage"`
	Knockback int64 `json:"knockback"`
}

type summonerWeapon struct {
	ManaCost  int64 `json:"mana_cost"`
	Damage    int64 `json:"damage"`
	Minion    int64 `json:"minion"`
	Knockback int64 `json:"knockback"`
}

type rougeWeapon struct {
}
