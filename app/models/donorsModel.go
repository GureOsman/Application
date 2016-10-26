package models
type Donors struct {
	ID             int     `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	ContactNo      int     `json:"contactno"`
	Gender         bool    `json:"gender"`
	Address        string  `json:"address"`
	BloodGroup     string  `json:"blood_group"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Age            int     `json:"age"`
}