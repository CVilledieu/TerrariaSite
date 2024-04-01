package main

import (
	info "Terraria/info"
	//server "Terraria/server"
	"fmt"
	"log"
)

func main() {
	//server.StartServer()
	newItem, err := info.GetToolInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("General Info", newItem.GeneralInfo)
	fmt.Println("General Info", newItem.GeneralInfo.ItemSpecs.Recipies)
	fmt.Println("Tool Speed", newItem.ToolSpeed)
	fmt.Println("Tool Damage", newItem.ToolDamage)
	fmt.Println("Power", newItem.Power)
	fmt.Println("Minable Ores", newItem.MinableOres)
}
