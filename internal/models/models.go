package models

type Embassy struct {
	HomeCountry string `bson:"home_country"`
	HostCountry string `bson:"host_country"`
	Name        string `bson:"name"`
	MapLink     string `bson:"map_link"`
	City        string `bson:"city"`
}
