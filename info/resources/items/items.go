package items

type ItemInfo struct {

	// Fields that apply to all items
	Name      string     `json:"name"`
	ItemSpecs *itemSpecs `json:"ItemSpecs"`  // ItemType refers to the type of item, e.g. weapon, armor, potion, etc.
	SellPrice int64      `json:"sell_price"` // Price refers to the price of the item when sold.
	StackSize int64      `json:"stack_size"` // StackSize refers to the maximum number of items that can be stacked together. Non-stackable items have a StackSize of 1.

}

type itemSpecs struct {
	Type string `json:"type"`

	WeaponInfo *weapon `json:"weapon_info"`
	ArmorInfo  *armor  `json:"armor_info"`
	PotionInfo *potion `json:"potion_info"`
	ToolInfo   *tool   `json:"tool_info"`

	Recipies   *recipe `json:"recipies"`
	ToolTip    string  `json:"tool_tip"`
	Difficulty string  `json:"difficulty"`
}

type recipe struct {
	UsedToCraft []string    `json:"used_to_craft"`
	CraftedFrom []*crafting `json:"crafted_from"`
	CraftedAt   string      `json:"crafted_at"`
}

type crafting struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}
