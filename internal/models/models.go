package models

type Embassy struct {
	//TODO: add bson tags
	HomeCountry string
	HostCountry string
	Name        string
	MapLink     string
	City        string
}
