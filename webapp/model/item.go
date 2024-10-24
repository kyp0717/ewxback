package model

import (
	"github.com/shopspring/decimal"
)

type Item struct {
//	SKU             string          `json:"SKU" gorm:"primarykey;type:text;size:20"`
	SKU             string          `json:"SKU" gorm:"column:SKU;type:text;size:20;null"`
	ItemName        string          `json:"ItemName" gorm:"column:ItemName;type:text;null"`
	UPC             string          `json:"UPC" gorm:"column:UPC;type:text;null"`
	Type            string          `json:"Type" gorm:"column:Type;type:text;null"`
	Category        string          `json:"Category" gorm:"column:Category;type:text;null"`
	Description     string          `json:"Description" gorm:"column:Description;type:text;null"`
	Inventory       int             `json:"Inventory" gorm:"column:Inventory;null"`
	QtyPerBox       int             `json:"QtyPerBox" gorm:"column:QtyPerBox;null"`
	Cost            decimal.Decimal `json:"Cost" gorm:"column:Cost;type:decimal(10,2);null"`
	Price           decimal.Decimal `json:"Price" gorm:"column:Price;type:decimal(10,2);null"`
	Price5          decimal.Decimal `json:"Price_5" gorm:"column:Price_5;type:decimal(10,2);null"`
	Price7          decimal.Decimal `json:"Price_7" gorm:"column:Price_7;type:decimal(10,2);null"`
	Price10         decimal.Decimal `json:"Price_10" gorm:"column:Price_10;type:decimal(10,2);null"`
	Price15         decimal.Decimal `json:"Price_15" gorm:"column:Price_15;type:decimal(10,2);null"`
	Price19         decimal.Decimal `json:"Price_19" gorm:"column:Price_19;type:decimal(10,2);null"`
	ItemDimension   string          `json:"ItemDimension" gorm:"column:ItemDimension;type:text;null"`
	Length          int             `json:"Length" gorm:"column:Length;null"`
	Width           int             `json:"Width" gorm:"column:Width;null"`
	Height          int             `json:"Height" gorm:"column:Height;null"`
	BoxDimension    string          `json:"BoxDimension" gorm:"column:BoxDimension;type:text;null"`
	BoxLength       int             `json:"Box_Length" gorm:"column:Box_Length;null"`
	BoxWidth        int             `json:"Box_Width" gorm:"column:Box_Width;null"`
	BoxHeight       int             `json:"Box_Height" gorm:"column:Box_Height;null"`
	BoxWeight       int             `json:"Box_Weight" gorm:"column:Box_Weight;null"`
	AvailableDate   string          `json:"AvailableDate" gorm:"column:AvailableDate;type:date;null"`
	ShippingMethod  string          `json:"Shipping_Method" gorm:"column:Shipping_Method;type:text;null"`
	PiecesContainer int             `json:"PiecesContainer" gorm:"column:PiecesContainer;null"`
	Supplier        string          `json:"Supplier" gorm:"column:Supplier;type:text;null"`
	ShippingCost    decimal.Decimal `json:"ShippingCost" gorm:"column:ShippingCost;type:decimal(10,2);null"`
	Active          string          `json:"Active" gorm:"column:Active;;type:text;size:1;null"`
	UserName        string          `json:"UserName" gorm:"column:UserName;type:text;size:20;null"`
	UpdateStamp     string          `json:"UpdateStamp" gorm:"column:UpdateStamp;type:date;null"`
}
