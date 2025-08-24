package models

type User struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
