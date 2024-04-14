package info

type FullItemInfo struct {
	Id               string
	Name             string
	Value            string
	UseStyle         string
	AnimationTime    string
	UseTime          string
	Consumable       bool
	UseTurn          bool
	AutoReuse        bool
	Rarity           string
	CreateTile       string
	Flavortext       string
	MaxStack         string
	HitboxWidth      string
	HitboxHeight     string
	PlaceStyle       string
	UseSound         string
	Damage           string
	Knockback        string
	ProjectileID     string
	ShootSpeed       string
	NoMeleeDamage    bool
	Accessory        bool
	DamageTypeMelee  bool
	DamageTypeMagic  bool
	DamageTypeRange  bool
	ToolTip2         string
	Defense          string
	Crit             string
	NoUseGraphic     string
	Scale            string
	CreateWall       string
	HeadSlot         string
	BodySlot         string
	LegSlot          string
	BuffType         string
	UseAmmo          string
	Ammo             string
	Mana             string
	Channel          bool
	BuffTime         string
	TileBoost        string
	NoWest           bool
	Glow             string
	AxePower         string
	PickAxePower     string
	MakeNPC          string
	Mech             string
	WingSlot         string
	Hammer           string
	HoldStyle        string
	Bait             string
	Expert           bool
	QuestItem        bool
	UniqueStack      string
	Paint            string
	HandOnSlot       string
	Flame            bool
	DamageTypeThrown bool
	Vanity           string
	Alpha            string
	ShoeSlot         string
	DamageTypeSummon bool
	StringColor      string
	BalloonSlot      string
	HandOffSlot      string
	MountType        string
	HealLife         string
	Potion           bool
	WaistSlot        string
	FishingPole      string
	BackSlot         string
	color            string
	NeckSlot         string
	FaceSlot         string
	HealMana         string
	Material         bool
	tileWand         string
	ShieldSlot       string
	NotAmmo          string
	FrontSlot        string
	CartTrack        bool
	ReuseDelay       string
	LifeRegen        string
	HairDye          string
	Stack            string
}

type Items interface {
	setWeaponValues()
	setArmorValues()
}

func setWeaponValues() {

}

func setArmorValues(id int, csvdb [][]string) *FullItemInfo {
	var item FullItemInfo
	item.Id = csvdb[id][0]
	item.Name = csvdb[id][1]

}
