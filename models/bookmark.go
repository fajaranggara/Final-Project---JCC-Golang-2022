package models

import "time"

type Bookmark struct {
	ID        int       `json:"id" gorm:"primary_key"`
	GameName  string    `json:"game_name"`
	IdGame		int		`json:"id_game"`
	Ratings		int		`json:"ratings"`
	ImageURL 	string 	`json:"image_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User	  User		`json:"-"`
}
