package models

import "time"

type Movie struct {
	ID          int       `json:"id" gorm:"primaryKey; autoIncrement:false"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Rating      int       `json:"rating"`
	MPAARating  string    `json:"mpaa_rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	MovieGenres []*Genre  `gorm:"many2many:movie_genres;"`
}

type Genre struct {
	ID        int       `json:"id" gorm:"primaryKey; autoIncrement:false"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Movies    []*Movie  `gorm:"many2many:movie_genres;"`
}
