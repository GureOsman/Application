package models
type Blood struct {
	ID           int       `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	Type         string    `json:"type"`
}
