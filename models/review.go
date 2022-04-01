package models

import "time"

type Review struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Rate  		int    	  `json:"rate"`
	Content		string    `json:"content"`
	GameID  	int		  `json:"game_id" gorm:"not null"`
	UserID		int		  `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Game		Game	  `json:"-"`
	User		User	  `json:"-"`
}