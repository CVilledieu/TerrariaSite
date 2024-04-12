package info

import (
	//items "Terraria/info/resources/items"
	//monsters "Terraria/info/resources/monsters"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	// ToolType is a constant that represents the type of item
	path = "info/resources/items/itemdb/"
)

type Item struct {
	Id         string
	Name       string
	Value      string
	FlavorText string
	StackSize  string
}

func ReadAllCSV() [][]string {
	file, err := os.Open("info/resources/items/items.csv")
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

func GetItem(name string) {
	record := ReadAllCSV()
	item := GetItemByName(name, record)
	fmt.Println(item)
}

func GetItemByName(name string, record [][]string) Item {
	foundItem := Item{}

	for i := 2; i < len(record); i++ {
		if record[i][1] == "name" {
			fmt.Println(record[i])
		}
	}

	return foundItem
}
