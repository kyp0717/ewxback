package model

type Item struct {
	Item string `json:"Item" gorm:"primaryKey";size:20"`

	SKU string `json:"SKU" gorm:"null;column:SKU;size:255"`
	UPC string `json:"UPC" gorm:"null;column:UPC;size:255"`
	Type string `json:"Type" gorm:"null;column:Type;size:255"`
	Category string `json:"Category" gorm:"null;column:Category;size:20"`
	Description string `json:"Description" gorm:"null;column:Category;size:255"`	
	Inventory Int `json:"Inventory" gorm:"null;column:Inventory"`
	Qty_per_Box Int `json:"Qty_per_Box" gorm:"null;column:Qty_per_Box"`
	Average_Cost float64 `json:"Average_Cost" gorm:"null;column:Average_Cost"`	
	Average_Price float64 `json:"Average_Price" gorm:"null;column:Average_Price"`		
}

