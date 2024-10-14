package model

type Test struct {
	Item string `json:"Item" gorm:"null";size:20"`
	Price float64 `json:"Price" gorm:"null;column:Price"`	
}

