package models

type User struct {
	ID       uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Email    string `json:"name" gorm:"unique"`
	Password string `json:"path"`
}
