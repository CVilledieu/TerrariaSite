package main

import (
	info "Terraria/info"
	"fmt"
	"log"
)

func main() {
	newItem, err := info.GetItemInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newItem)
}
