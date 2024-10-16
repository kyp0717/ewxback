package dataload

type Item struct {
	SKU              string  `json:"SKU" gorm:"column:SKU;type:text"`
	ItemName         string  `json:"Item" gorm:"column:Item;type:text"`
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

func item_load() {
	// Initialize DB connection
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto migrate the Item table
	err = db.AutoMigrate(&Item{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	// Load CSV data
	csvFile, err := os.Open("items.csv") // Replace with your CSV file path
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Read all rows from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Loop through the records and create Items
	for _, record := range records {
		item := Item{
			SKU:             record[0],
			ItemName:        record[1],
			UPC:             record[2],
			Type:            record[3],
			Category:        record[4],
			Description:     record[5],
			Inventory:       parseInt(record[6]),
			QtyPerBox:       parseInt(record[7]),
			AverageCost:     parseFloat(record[8]),
			AveragePrice:    parseFloat(record[9]),
			Dims:            record[10],
			BDims:           record[11],
			Length:          parseInt(record[12]),
			Width:           parseInt(record[13]),
			Height:          parseInt(record[14]),
			BoxDims:         record[15],
			BoxLength:       parseInt(record[16]),
			BoxWidth:        parseInt(record[17]),
			BoxHeight:       parseInt(record[18]),
			BoxWeight:       parseInt(record[19]),
			AvailableDate:   record[20],
			ShippingMethod:  record[21],
			PiecesContainer: parseInt(record[22]),
			YTDInventory:    parseInt(record[23]),
			Supplier:        record[24],
			UpdateStamp:     record[25],
		}

		// Save item to the database
		err = db.Create(&item).Error
		if err != nil {
			log.Printf("Failed to insert record: %v", err)
		}
	}

	fmt.Println("Data imported successfully!")
}

// Helper function to parse integers
func parseInt(value string) int {
	if value == "" {
		return 0
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}

// Helper function to parse floats
func parseFloat(value string) float64 {
	if value == "" {
		return 0.0
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return v
}
