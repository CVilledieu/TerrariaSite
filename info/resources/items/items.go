package items

type GeneralInfo struct {
	// Fields that apply to all items
	Id            int64  `json:"id"` // Id refers to the unique identifier of the item.
	Name          string `json:"name"`
	SellPrice     int64  `json:"sell_price"` // In copper coins. 100 copper coins = 1 silver coin and 100 silver coins = 1 gold coin 100 gold coins = 1 platinum coin
	StackSize     int64  `json:"stack_size"` // StackSize refers to the maximum number of items that can be stacked together. Non-stackable items have a StackSize of 1.
	UseTime       int64  `json:"use_time"`   // The number of frames before the item can be used again.
	Type          string `json:"type"`
	SecondaryType string `json:"secondary_type"`
	ToolTip       string `json:"tool_tip"`
	Difficulty    string `json:"difficulty"`
	Recipies      Recipe `json:"recipies"`
}

type Recipe struct {
	UsedToCraft []Crafting `json:"used_to_craft"`
	CraftedFrom []Crafting `json:"crafted_from"`
	CraftedAt   []Crafting `json:"crafted_at"`
}

type Crafting struct {
	Type   string `json:"type"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type Testinfo struct {
	// Fields that apply to all items
	Name      string `json:"name"`
	SellPrice int64  `json:"sell_price"` //In copper coins. 100 copper coins = 1 silver coin and 100 silver coins = 1 gold coin 100 gold coins = 1 platinum coin
	StackSize int64  `json:"stack_size"` // StackSize refers to the maximum number of items that can be stacked together. Non-stackable items have a StackSize of 1.
}
