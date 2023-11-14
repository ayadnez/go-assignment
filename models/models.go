package models

import (
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name`
	Mobile     string    `json:"mobile`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at`
}

type Product struct {
	ProductID               int       `json:"product_id"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImage            []string  `json:"product_images" gorm:"type:text"`
	ProductPrice            float64   `json:"product_price"`
	CompressedProductImages []string  `json:"compressed_images" gorm:"type:text"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type DbProduct struct {
	ProductID               int       `json:"product_id" gorm:"primaryKey"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImage            string    `json:"product_images" gorm:"type:text"`
	ProductPrice            float64   `json:"product_price"`
	CompressedProductImages string    `json:"compressed_images" gorm:"type:text"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func ToString(arr []string) string {
	result := strings.Join(arr, ",")

	return result

}

func ToSlice(str string) []string {
	return strings.Split(str, ",")

}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=zahid password=1234@123 dbname=db1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Open a connection to the database
	if err != nil {
		return nil, err
	}

	// AutoMigrate will attempt to automatically migrate your schema, creating the Product table
	err = db.AutoMigrate(&DbProduct{}, &User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
