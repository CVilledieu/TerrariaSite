package main

import (
	"Terraria/info"
	"fmt"
)

func main() {
	//startServer()
	items := info.LoadAllItems()
	fmt.Println(items)
}

func searchResults(name string) {
	//info.GetItem()
}
