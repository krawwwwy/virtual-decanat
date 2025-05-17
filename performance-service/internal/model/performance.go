package model

type StudentPerformance struct {
	StudentID   string
	Grades      []Grade
	Attendance  []Attendance
	Debts       []Debt
	Rating      float64
}

type Grade struct {
	Discipline string
	Value      float64
}

type Attendance struct {
	Discipline string
	Total      int
	Attended   int
}

type Debt struct {
	Discipline string
	Reason     string
} 