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

	//get house by id
	func GetHouseById(houseId int) *models.House {
		var house models.House
		db := DB()

		if db != nil {
			db.Where("id = ?", houseId ).First(&house)
			return &house
		}
		return nil
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
  // validar qtd de membros com aquela casa
func DeleteHouseById(houseId int) int {
	db := DB()
	var house models.House

	result := db.Where("id = ?",houseId).Delete(&house)

	if result.RowsAffected == 1 {
		return 1
	}
	return -1
}


//members helpers

func GetQuantityPersonByHouseID(houseId int) int{
	db := DB()
	var persons []models.Person

	if db != nil {
		result:=db.Where("houseid = ?", houseId).Find(&persons)
		return int(result.RowsAffected)
	}
	
	return -1
}


// get all persons
func GetAllPersons() *[]models.Person {
	var person []models.Person

	db := DB()

	if db != nil {
		db.Find(&person)
	}
	
	return &person
  }

  // get person by id
  func GetPersonById(personId int) *models.Person {
	var person models.Person

	db := DB()

	if db != nil {
		db.Where("id = ?", personId).Find(&person)
	}
	
	return &person
  }

// get all person
func GetPersonsByHouseId(houseId int) *[]models.Person {
	var persons []models.Person

	db := DB()

	if db != nil {
		db.Model(&persons).Select("id, name, is_married").Find(&persons)
	}
	
	return &persons
  }

    // add new house to the DB
	func AddNewPerson(person models.Person) int {
		db := DB()
	
		result := db.Create(&person)
	
		if result.RowsAffected == 0 || result.Error != nil {
			return -1
		}
	
		return 1
	
	  }

	// delete person by id 
	func DeletePersonById(personId int) int {
		db := DB()
		var person models.Person

		result := db.Where("id = ?", personId).Delete(&person)
		if result.RowsAffected > -1 {
			return int(result.RowsAffected)
		}
		return -1
	}
