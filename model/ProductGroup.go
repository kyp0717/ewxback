package model

type ProductGroup struct {
	SKU        string `json:"SKU" gorm:"column:SKU;type:text"`
	TotalBoxes int    `json:"Total_Boxes" gorm:"column:Total_Boxes;null"`
	ItemsSKU   string `json:"Items_SKU" gorm:"column:Items_SKU;type:text"`
	Box        int    `json:"BOX" gorm:"column:BOX;null"`
	Piece      int    `json:"Piece" gorm:"column:Piece;null"`
}
