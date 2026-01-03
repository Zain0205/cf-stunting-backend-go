// Package database
package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error

	for i := 1; i <= 10; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connected")
			return
		}

		fmt.Printf("Waiting for MySQL... (%d/10)\n", i)
		time.Sleep(3 * time.Second)
	}

	panic("Failed to connect MySQL after retries")
}
