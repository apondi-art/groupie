package models

type Apis struct {
	ApisUrls map[string]string
}

func ConfigApis() *Apis {
	var urls = make(map[string]string)

	urls["artists"] = "https://groupietrackers.herokuapp.com/api/artists"
	urls["locations"] = "https://groupietrackers.herokuapp.com/api/locations"
	urls["dates"] = "https://groupietrackers.herokuapp.com/api/dates"
	urls["relations"] = "https://groupietrackers.herokuapp.com/api/relation"

	return &Apis{urls}
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
