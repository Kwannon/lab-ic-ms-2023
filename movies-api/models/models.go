package models

type Director struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Movie struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Description string `json:"description" gorm:"text"`
	DirectorID  int    `json:"-"`
	Director    Director
}
