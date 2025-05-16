package models

type Application struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Surname   string   `json:"surname"`
	Otchestvo string   `json:"otchestvo"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Number    string   `json:"number"`
	Faculty   string   `json:"faculty"`
	Snils     string   `json:"snils"`
	Passport  Passport `json:"passport"`
	Exams     Exams    `json:"exams"`
	Status    string   `json:"status"`
}
