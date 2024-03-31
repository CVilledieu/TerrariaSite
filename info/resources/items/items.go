package items

type ItemInfo struct {

	// Fields that apply to all items
	Name      string       `json:"name"`
	ItemSpecs []*ItemSpecs `json:"ItemSpecs"`  // ItemType refers to the type of item, e.g. weapon, armor, potion, etc.
	SellPrice int64        `json:"sell_price"` // Price refers to the price of the item when sold.
	StackSize int64        `json:"stack_size"` // StackSize refers to the maximum number of items that can be stacked together. Non-stackable items have a StackSize of 1.

}

type ItemSpecs struct {
	ItemType   string    `json:"item_type"`
	WeaponInfo []*weapon `json:"weapon_info"`
	ArmorInfo  []*armor  `json:"armor_info"`
	PotionInfo *potion   `json:"potion_info"`
	ToolInfo   []*tool   `json:"tool_info"`
	Recipies   []*recipe `json:"recipies"`
}

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
}

type rangedWeapon struct {
}

type magicWeapon struct {
}

type summonerWeapon struct {
}

type rougeWeapon struct {
}

// Armor items
type armor struct {
	ArmorType string  `json:"armor_type"`
	Helmet    *helmet `json:"helmet"`
	Shirt     *shirt  `json:"shirt"`
	Pants     *pants  `json:"pants"`
}

type helmet struct {
	Class   string `json:"class"`
	Defense int64  `json:"defense"`
	Effect  string `json:"effect"`
}

type shirt struct {
	Defense int64  `json:"defense"`
	Effect  string `json:"effect"`
}

type pants struct {
	Defense int64  `json:"defense"`
	Effect  string `json:"effect"`
}

type accessory struct {
}

type pets struct {
}
type mount struct {
}

type cosmetic struct {
}

// consumable items
type potion struct {
}

// Crafting recipes
type recipe struct {
	CraftedInto []string `json:"crafted_into"`
	CraftedFrom []string `json:"crafted_from"`
	CraftedAt   []string `json:"crafted_at"`
}

// Miscellaneous items
type key struct {
}

type bossSummon struct {
}

// crafting materials
type OreBar struct {
}

type OreRaw struct {
}

type monsterDrop struct {
}

type plant struct {
}

// Placeable items
type craftingBenches struct {
}

type furniture struct {
}
