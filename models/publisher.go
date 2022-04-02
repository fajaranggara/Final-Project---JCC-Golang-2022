package models

import "time"

type Publisher struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	ImageURL 	string    `json:"image_url"`
	UserID		int		  `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Game		[]Game	  `json:"-"`
	User		User	  `json:"-"`
}
