package model

type Product struct {
	SKU         string `json:"SKU" gorm:"primarykey;type:text"`
	Product     string `json:"Product" gorm:"column:Product;type:text"`
	CreateDate  string `json:"Create_Date" gorm:"column:Create_Date;type:text"`
	TotalBoxes  int    `json:"Total_Boxes" gorm:"column:Total_Boxes;null"`

	SKU1        string `json:"SKU_1" gorm:"column:SKU_1;type:text"`
	Box1        int    `json:"BOX_1" gorm:"column:BOX_1;null"`
	Piece1      int    `json:"Piece_1" gorm:"column:Piece_1;null"`

	SKU2        string `json:"SKU_2" gorm:"column:SKU_2;type:text"`
	Box2        int    `json:"BOX_2" gorm:"column:BOX_2;null"`
	Piece2      int    `json:"Piece_2" gorm:"column:Piece_2;null"`

	SKU3        string `json:"SKU_3" gorm:"column:SKU_3;type:text"`
	Box3        int    `json:"BOX_3" gorm:"column:BOX_3;null"`
	Piece3      int    `json:"Piece_3" gorm:"column:Piece_3;null"`

	SKU4        string `json:"SKU_4" gorm:"column:SKU_4;type:text"`
	Box4        int    `json:"BOX_4" gorm:"column:BOX_4;null"`
	Piece4      int    `json:"Piece_4" gorm:"column:Piece_4;null"`

	SKU5        string `json:"SKU_5" gorm:"column:SKU_5;type:text"`
	Box5        int    `json:"BOX_5" gorm:"column:BOX_5;null"`
	Piece5      int    `json:"Piece_5" gorm:"column:Piece_5;null"`

	SKU6        string `json:"SKU_6" gorm:"column:SKU_6;type:text"`
	Box6        int    `json:"BOX_6" gorm:"column:BOX_6;null"`
	Piece6      int    `json:"Piece_6" gorm:"column:Piece_6;null"`

	SKU7        string `json:"SKU_7" gorm:"column:SKU_7;type:text"`
	Box7        int    `json:"BOX_7" gorm:"column:BOX_7;null"`
	Piece7      int    `json:"Piece_7" gorm:"column:Piece_7;null"`
}
