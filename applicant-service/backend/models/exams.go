package models

type Exams struct {
	Russian     int `json:"russian"`
	Math        int `json:"math"`
	Physics     int `json:"physics"`
	Informatics int `json:"informatics"`
}
