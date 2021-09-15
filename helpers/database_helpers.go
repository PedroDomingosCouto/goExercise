package helpers

import (
	"fmt"
	"goExercise/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase(){
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf( "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db_name, port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Error connecting to the database - %s", err.Error())
	}
}

  func DB() *gorm.DB{
	return db
  }

// get all houses
  func GetAllHouses() *[]models.House {
	var houses []models.House

	db := DB()

	if db != nil {
		db.Find(&houses)
	}
	
	return &houses
  }

  // add new house to the DB
  func CreateHouse(house models.House) int {
	db := DB()

	result := db.Create(&house)

	if result.RowsAffected == 0 || result.Error != nil {
		return -1
	}

	return 1

  }

  // update housa data

  func UpdateHouseData() int {
	  return 1
  }


  // delete house by id 

  func RemoveHouseById() int {
	  return 1
  }