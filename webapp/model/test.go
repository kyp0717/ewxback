package model
import (
  "github.com/shopspring/decimal"
)
type Test1 struct {
  Item string   `json:"item"  gorm:"null;column:item;size:200"`
  // Price float64 `json:"price" gorm:"null;column:price;size:255"`	
  Price decimal.Decimal `json:"price" gorm:"null;column:price;type:decimal(20,8);"`
}



type Test2 struct {
  Item1 string   `json:"item1"  gorm:"null;column:item1;size:200"`
  // Price float64 `json:"price" gorm:"null;column:price;size:255"`	
  Item2 string   `json:"item2"  gorm:"null;column:item2;size:200"`
}

