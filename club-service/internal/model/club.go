package model

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

// Club represents a club entity
type Club struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
} 