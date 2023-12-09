package main

import "fmt"

// anonymous structure
type Player struct {
	string
	int
}

type Hero struct {
	name string
	role string
}

func main() {
	// initialize()
	initialize2()
}

func initialize() {
	var player = Player{"Agung", 21}
	fmt.Println(player)
	
	var hero = Hero{"Adith", "Tank"}
	fmt.Println(hero)
}

func initialize2() {
	var player = new (Player)
	player.string = "Agung"
	player.int = 21
	fmt.Println(player)
	
	var hero = new (Hero)
	hero.name = "Adith"
	hero.role = "Tank"
	fmt.Println(hero)
}
