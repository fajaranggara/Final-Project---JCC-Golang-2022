package models

import "time"

type Publisher struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	ImageURL 	string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Game		[]Game	  `json:"-"`
}
