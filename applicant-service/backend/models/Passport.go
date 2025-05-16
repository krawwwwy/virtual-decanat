package models

type Passport struct {
	Seial  int    `json:"seial"`
	Number int    `json:"number"`
	Date   string `json:"date"`
	Adress string `json:"adress"`
}
