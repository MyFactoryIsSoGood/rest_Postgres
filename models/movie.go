package models

type Movie struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	ReleaseYear uint16 `json:"release_year"`
}
