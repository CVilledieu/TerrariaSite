package info

import (
	items "Terraria/info/resources/items"
	monsters "Terraria/info/resources/monsters"
	"encoding/json"
	"os"
)

func GetBossInfo(name string) (string, error) {
	var monster monsters.MonsterInfo
	monster.Name = name
	return monster.Name, nil
}

func GetItemInfo() (*items.ItemInfo, error) {
	jsonFile, _ := os.ReadFile("./items/itemdb/tools.json")
	var item items.ItemInfo
	err := json.Unmarshal(jsonFile, &item) //
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func GetAllItems() ([]string, error) {

	return nil, nil
}
