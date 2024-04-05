package main

import (
	info "Terraria/info"
	server "Terraria/server"

	"log"
)

func main() {
	server.StartServer()
	item, err := info.GetToolInfo("Tin Pickaxe")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(item)
}
