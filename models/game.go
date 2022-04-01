package models

import "time"

type Game struct {
	ID          int       	`json:"id" gorm:"primary_key"`
	Name        string    	`json:"name" gorm:"not null"`
	Ratings		int			`json:"ratings"`
	RatingsCounter int      `json:"ratings_counter"`		
	ReleaseDate string    	`json:"release_date"`
	Description string 		`json:"description"`
	ImageURL 	string 		`json:"image_url"`
	GenreID		int			`json:"genre_id"`
	CategoryID	int			`json:"category_id"`
	PublisherID	int			`json:"publisher_id"`
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
	Genre		Genre		`json:"-"`
	Category	Category	`json:"-"`
	Publisher	Publisher	`json:"-"`
	Review		[]Review	`json:"-"`
}
