package models

type Recipe struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Ingredients string `json:"ingredients"`
	Type        string `json:"type"` // entrée, plat, dessert
	IsVegan     bool   `json:"is_vegan"`
}
