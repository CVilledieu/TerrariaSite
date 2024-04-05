package info

import (
	items "Terraria/info/resources/items"
	monsters "Terraria/info/resources/monsters"
	"encoding/json"
	"os"
)

const (
	// ToolType is a constant that represents the type of item
	path = "info/resources/items/itemdb/"
)

func GetBossInfo(bossname string) (string, error) {
	var monster monsters.MonsterInfo
	//monster.Name = name
	return monster.Name, nil
}

func GetToolInfo(toolname string) (*items.Tool, error) {

	jsonFile, _ := os.ReadFile(path + "tools.json")
	var item items.Tool
	err := json.Unmarshal(jsonFile, &item) //
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func GetWeaponInfo(weaponname string) (*items.Weapon, error) {
	jsonFile, _ := os.ReadFile("./items/itemdb/weapons.json")
	var item items.Weapon
	err := json.Unmarshal(jsonFile, &item) //
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func GetArmorInfo(armorname string) (*items.Armor, error) {
	jsonFile, _ := os.ReadFile("./items/itemdb/armor.json")
	var item items.Armor
	err := json.Unmarshal(jsonFile, &item) //
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func GetAllItems() ([]string, error) {

	return nil, nil
}

func GetTestInfo() (*items.Testinfo, error) {
	jsonFile, err := os.ReadFile(path + "test.json")
	if err != nil {
		return nil, err
	}
	var item items.Testinfo
	err2 := json.Unmarshal(jsonFile, &item) //
	if err2 != nil {
		return nil, err2
	}

	return &item, nil

}
