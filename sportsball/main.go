package main

import "fmt"

type Player struct {
	firstName string
	lastName  string
	number    int
}

type Roster struct {
	players []Player
}

type Team struct {
	mascot   string
	location string
	roster   Roster
}

func main() {
	var rockets Team
	var roster Roster
	roster.players = append(roster.players, Player{"James", "Harden", 32})
	roster.players = append(roster.players, Player{"CJ", "mcullum", 15})
	rockets.roster = roster
	fmt.Println(rockets)
}
