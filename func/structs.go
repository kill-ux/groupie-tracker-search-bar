package groupie

// is the data of Relations for one artist
type Concert struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// is the data of ConcertDates for one artist
type DataConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// is the data of Locations for one artist
type DataLocations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:dates`
}

// data of one artist
type Artist struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	ConcertDates     string   `json:"concertDates"`
	Relations        string   `json:"relations"`
	DataLocations    DataLocations
	DataConcertDates DataConcertDates
	Concerts         Concert
}

// data to be shown in the page
type Page struct {
	Code       int
	MsgError   string
	Arts       []Artist
	Art        Artist
	Filters    []Artist
	DataLocals map[string][]DataLocations
	DataList   map[any]string
	// FromCreationDate string
	// ToCreationDate   string
	// FromFirsetAlbum  string
	// ToFirsetAlbum    string
	// Local            string
	// Members          []string
}
