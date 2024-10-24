package model

type Item struct {
	SKU              string  `json:"SKU" gorm:"primarykey;type:text"`
	ItemName       	 string  `json:"ItemName" gorm:"column:ItemName;type:text"`	
	UPC              string  `json:"UPC" gorm:"column:UPC;type:text"`
	Type             string  `json:"Type" gorm:"column:Type;type:text"`
	Category         string  `json:"Category" gorm:"column:Category;type:text"`
	Description      string  `json:"Description" gorm:"column:Description;type:text"`
	Inventory        int     `json:"Inventory" gorm:"column:Inventory;null"`
	QtyPerBox        int     `json:"Qty_per_Box" gorm:"column:Qty_per_Box;null"`
	AverageCost      float64 `json:"Average_Cost" gorm:"column:Average_Cost;type:decimal(10,2);null"`
	AveragePrice     float64 `json:"Average_Price" gorm:"column:Average_Price;type:decimal(10,2);null"`
	Dims             string  `json:"Dims" gorm:"column:Dims;type:text"`
	BDims            string  `json:"B_Dims" gorm:"column:B_Dims;type:text"`
	Length           int     `json:"Length" gorm:"column:Length;null"`
	Width            int     `json:"Width" gorm:"column:Width;null"`
	Height           int     `json:"Height" gorm:"column:Height;null"`
	BoxDims          string  `json:"Box_Dims" gorm:"column:Box_Dims;type:text"`
	BoxLength        int     `json:"Box_Length" gorm:"column:Box_Length;null"`
	BoxWidth         int     `json:"Box_Width" gorm:"column:Box_Width;null"`
	BoxHeight        int     `json:"Box_Height" gorm:"column:Box_Height;null"`
	BoxWeight        int     `json:"Box_Weight" gorm:"column:Box_Weight;null"`
	AvailableDate    string  `json:"Available_Date" gorm:"column:Available_Date;type:date;null"`
	ShippingMethod   string  `json:"Shipping_Method" gorm:"column:Shipping_Method;type:text"`
	PiecesContainer  int     `json:"Pieces_Container" gorm:"column:Pieces_Container;null"`
	YTDInventory     int     `json:"YTD_Inventory" gorm:"column:YTD_Inventory;null"`
	Supplier         string  `json:"Supplier" gorm:"column:Supplier;type:text"`
	UpdateStamp      string  `json:"UpdateStamp" gorm:"column:UpdateStamp;type:date;null"`
}
