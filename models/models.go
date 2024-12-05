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
