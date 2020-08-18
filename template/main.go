package main

import (
	"html/template"
	"log"
	"os"
)

type team struct {
	City string
	Name string
	Seed int8
}

type matchup struct {
	T1     team
	T2     team
	T1Wins int8
	T2Wins int8
}

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	//West
	lakers := team{"Los Angeles", "Lakers", 1}
	blazers := team{"Portland", "Trail Blazers", 8}
	rockets := team{"Houston", "Rockets", 4}
	thunder := team{"Oklahoma City", "Thunder", 5}
	nuggets := team{"Denver", "Nuggets", 3}
	jazz := team{"Utah", "Jazz", 6}
	clippers := team{"Los Angeles", "Clippers", 2}
	mavs := team{"Dallas", "Mavericks", 7}

	//East
	bucks := team{"Milwaukee", "Bucks", 1}
	magic := team{"Orlando", "Magic", 2}
	pacers := team{"Indiana", "Pacers", 4}
	heat := team{"Miami", "Heat", 5}
	celtics := team{"Boston", "Celtics", 3}
	sixers := team{"Phillidelphia", "76ers", 6}
	raptors := team{"Toronto", "Raptors", 2}
	nets := team{"Brooklyn", "Nets", 7}

	//First Round
	matchups := []matchup{
		{T1: lakers, T2: blazers},
		{T1: rockets, T2: thunder},
		{T1: nuggets, T2: jazz},
		{T1: clippers, T2: mavs},
		{T1: bucks, T2: magic},
		{T1: pacers, T2: heat},
		{T1: celtics, T2: sixers},
		{T1: raptors, T2: nets},
	}

	//Semifinals
	//Conference Championship
	//Finals

	f, err := os.Create("index.html")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, matchups)
	if err != nil {
		log.Fatal(err)
	}
}
