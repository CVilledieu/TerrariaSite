package items

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
