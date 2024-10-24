package data

// Item model definition
import 
(
	"github.com/kyp0717/ewxback/webapp/model"
	"fmt"
	"log"
   	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func item_select() {
	// Initialize DB connection
	dsn := "host=localhost user=youruser password=password dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Query all items
	var items []model.Item
		err = db.Find(&items).Error
	if err != nil {
		log.Fatalf("Error retrieving items: %v", err)
	}

	// Display the result
	for _, item := range items {
		fmt.Printf("SKU: %s, Item: %s, UPC: %s, Inventory: %d, Price: %.2f\n", item.SKU, item.ItemName, item.UPC, item.Inventory, item.Price)
	}
}
