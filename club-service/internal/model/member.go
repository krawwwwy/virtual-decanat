package model

// Member represents a club member entity
type Member struct {
	ID       int    `json:"id"`
	ClubID   int    `json:"club_id"`
	UserID   int    `json:"user_id"`
	Role     string `json:"role"`
	JoinedAt string `json:"joined_at"`
} 