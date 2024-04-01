package main

import (
	info "Terraria/info"
	//server "Terraria/server"
	"fmt"
	"log"
)

func main() {
	//server.StartServer()
	newItem, err := info.GetItemInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newItem)
}
