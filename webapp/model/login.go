package model

type Login struct {
	UserName string`json:"Username" gorm:"primaryKey"`
	Password string `Password:"title" gorm:"not null;column:Password;size:25"`
}
