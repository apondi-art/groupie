package models

type Apis struct {
	Artist    string
	Locations string
	Date      string
	Relation  string
}

func NewApis() *Apis {
	return &Apis{
		Artist:    "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Date:      "https://groupietrackers.herokuapp.com/api/dates",
		Relation:  "https://groupietrackers.herokuapp.com/api/relation",
	}
}

type Artist struct {
	ConcertDates string   `json:"concertDates"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Locations    string   `json:"locations"`
	Members      []string `json:"members"`
	Relations    string   `json:"relations"`
}
