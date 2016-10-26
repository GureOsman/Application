package models
type City struct {
	ID           int       `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	Name         string    `json:"name"`
}
