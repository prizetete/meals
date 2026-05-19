package main

import (
	"fmt"
	"komkrit/adapters"
	"komkrit/core"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ├── adapters --> สำหรับเก็บ adapter
// │   ├── gorm_adapter.go
// │   └── http_adapter.go
// ├── core --> สำหรับเก็บ business logic, port
// │   ├── order.go
// │   ├── order_repository.go
// │   └── order_service.go
// ├── main.go

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// Initialize the database connection
	// db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	app := fiber.New()

	// Set up the core service and adapters
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// Define routes
	app.Post("/order", orderHandler.SSssss)

	app.Listen(":8080")
}
