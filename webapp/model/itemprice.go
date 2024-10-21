package model

type ItemPrice struct {
	SKU           string  `json:"SKU" gorm:"primarykey;type:text"`

	Cost          float64 `json:"Cost" gorm:"column:Cost;type:decimal(10,2);null"`
	ShippingCost  float64 `json:"Shipping_Cost" gorm:"column:Shipping_Cost;type:decimal(10,2);null"`
	Price         float64 `json:"Price" gorm:"column:Price;type:decimal(10,2);null"`
	Price5        float64 `json:"Price_5" gorm:"column:Price_5;type:decimal(10,2);null"`
	Price7        float64 `json:"Price_7" gorm:"column:Price_7;type:decimal(10,2);null"`
	Price10       float64 `json:"Price_10" gorm:"column:Price_10;type:decimal(10,2);null"`
	Price15       float64 `json:"Price_15" gorm:"column:Price_15;type:decimal(10,2);null"`
	Price19       float64 `json:"Price_19" gorm:"column:Price_19;type:decimal(10,2);null"`
	Inventory     int     `json:"Inventory" gorm:"column:Inventory;null"`
	Active        string  `json:"Active" gorm:"column:Active;type:varchar(1);null"`
	UpdateStamp   string  `json:"UpdateStamp" gorm:"column:UpdateStamp;type:date;null"`
}
