package main

import (
	"sync"

	Groupie "groupie/func"
)

func init() {
	// Fetch a data
	var wg sync.WaitGroup
	wg.Add(2)
	go Groupie.Fetch(&wg, "https://groupietrackers.herokuapp.com/api/artists", &Groupie.Data.Arts)
	go Groupie.Fetch(&wg, "https://groupietrackers.herokuapp.com/api/locations", &Groupie.Data.DataLocals)
	wg.Wait()
}
