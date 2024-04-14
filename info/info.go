package info

import (
	//items "Terraria/info/resources/items"
	//monsters "Terraria/info/resources/monsters"
	"encoding/csv"
	"log"
	"os"
)

type Item struct {
	Id          string
	Name        string
	OtherValues [][]string
}

func ReadAllCSV() [][]string {
	file, err := os.Open("info/resources/items.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)

	record, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return record
}

func GetItem(name string) Item {
	record := ReadAllCSV()
	item := getItemByName(name, record)
	return item
}

func getItemByName(name string, record [][]string) Item {
	foundItem := Item{Id: name}
	for i := 2; i < len(record); i++ {
		if record[i][1] == name {
			foundItem.Name = record[i][0]
			foundItem.OtherValues = getItemInfo(i, record)
			break
		}
	}
	return foundItem
}

func getItemInfo(index int, record [][]string) [][]string {
	var itemInfo [][]string
	for i, attribute := range record[index] {
		if attribute != "-" {
			itemInfo = append(itemInfo, []string{record[0][i], attribute})
		}
	}
	return itemInfo
}
