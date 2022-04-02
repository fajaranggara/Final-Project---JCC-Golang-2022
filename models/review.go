package models

import "time"

type Review struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Rate  		int    	  `json:"rate" gorm:"not null, min=1, max=5"`
	Content		string    `json:"content"`
	GameID  	int		  `json:"game_id"`
	UserID		int		  `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Game		Game	  `json:"-"`
	User		User	  `json:"-"`
}
