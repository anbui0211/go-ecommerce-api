package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   int
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "An"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True" // &local=Local

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Error dropping tables: ", err)
		}
	}

	// Tạo bảng nếu chưa có
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập các tham số kết nối
	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:33306)/shopdevgo?charset=utf8mb4&parseTime=True" // &local=Local
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Error dropping tables: ", err)
		}
	}

	// Tạo bảng nếu chưa có
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập các tham số kết nối
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			insertRecord(b, db)
		}
	})
}

// go test -bench=. -benchmem
