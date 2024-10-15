package test

import (
	"github.com/kyp0717/ewxback/model"
  "os"
  "log"
  "fmt"
  "encoding/csv"
	"gorm.io/gorm"
  "strconv"
) 

func Testload(db *gorm.DB ) {

	// Load CSV data
	csvFile, err := os.Open("test/test.csv") // Replace with your CSV file path
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Read all rows from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV test file: %v", err)
	}

	// Loop through the records and create Items
	for _, record := range records {
    price, _ := strconv.ParseFloat(record[1], 64)
		test := &model.Test{
			Item:         record[0],
			Price:        price,
		}

		// Save item to the database
		err = db.Create(&test).Error
		if err != nil {
			log.Printf("Failed to insert test record: %v", err)
		}
	}

	fmt.Println("Test Data imported successfully!")
}

