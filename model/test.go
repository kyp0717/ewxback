package model
import (
  "github.com/shopspring/decimal"
)
type TestTable struct {
  Item string   `json:"item"  gorm:"null;column:item;size:200"`
  // Price float64 `json:"price" gorm:"null;column:price;size:255"`	
  Price decimal.Decimal `json:"amount" gorm:"null;column:price;type:decimal(20,8);"`
}

