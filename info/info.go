package info

import (
	items "Terraria/info/resources/items"
	monsters "Terraria/info/resources/monsters"
)

func GetBossInfo(name string) (string, error) {
	var monster monsters.MonsterInfo
	monster.Name = name
	return monster.Name, nil
}

func GetItemInfo() (*items.ItemInfo, error) {
	var item items.ItemInfo
	item.Name = "Sword"
	item.SellPrice = 10
	item.StackSize = 1
	return &item, nil

}

func GetAllItems() ([]string, error) {

	return nil, nil
}
