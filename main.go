package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie "groupie/func"
)

func main() {
	if len(Groupie.Data.Arts) == 0 {
		log.Fatalf("Error fetching data")
	}
	http.HandleFunc("/", Groupie.HandelHome)
	http.HandleFunc("/artist/", Groupie.HandelArtist)
	http.HandleFunc("/artist", Groupie.HandelArtist)
	http.HandleFunc("/css/", Groupie.CssHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
