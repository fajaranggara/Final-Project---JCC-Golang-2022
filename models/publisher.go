package models

import "time"

type Publisher struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	LogoURL 	string    `json:"logo_url"`
	UserID		int		  `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Game		[]Game	  `json:"-"`
	User		User	  `json:"-"`
}
