package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// to show the dtails of specific artist
func HandelArtist(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}

	id := strings.TrimPrefix(req.URL.Path, "/artist/")
	idTemp, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	if idTemp < 1 || idTemp > len(Data.Arts) {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	Data.Art = Data.Arts[idTemp-1]
	var wg sync.WaitGroup
	wg.Add(3)
	go Fetch(&wg, Data.Art.Relations, &Data.Art.Concerts)
	go Fetch(&wg, Data.Art.Locations, &Data.Art.DataLocations)
	go Fetch(&wg, Data.Art.ConcertDates, &Data.Art.DataConcertDates)
	wg.Wait()
	RenderPage("artist", res)
}

// Fetch a data
func Fetch(wg *sync.WaitGroup, url string, data any) {
	defer wg.Done()
	resGet, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err.Error())
	}
	defer resGet.Body.Close()
	if resGet.StatusCode != http.StatusOK {
		log.Fatalf("Error: Status code %d ", resGet.StatusCode)
	}
	if err := json.NewDecoder(resGet.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}
