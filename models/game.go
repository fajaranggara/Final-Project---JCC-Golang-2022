package models

import "time"

type Game struct {
	ID          int       	`json:"id" gorm:"primary_key"`
	Name        string    	`json:"name" gorm:"not null"`
	Ratings		float64		`json:"ratings"`		
	ReleaseDate string    	`json:"release_date"`
	Price 		int 		`json:"price"`
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
