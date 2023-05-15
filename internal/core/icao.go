package core

type Airlines struct {
	Airline     string `csv:"Airline"`
	Callsign    string `csv:"Callsign"`
	Name        string `csv:"Name"`
	CountryCode string `csv:"Country code"`
}
