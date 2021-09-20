package helpers

import "goExercise/models"

// get all persons
func GetAllPersons() *[]models.Person {
	var person []models.Person

	db := DB()

	if db != nil {
		db.Find(&person)
		return &person
	}
	
	return nil
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
		db.Model(&persons).Select("id, name, is_married").Where("house_id= ?", houseId).Find(&persons)
		return &persons
	}

	return nil
  }

    // add new house to the DB
	func AddNewPerson(person models.Person) int {
		db := DB()
		
		if db != nil {
			result := db.Create(&person)
			if result.RowsAffected == 0 || result.Error != nil {
				return -1
			}
			return 1
		}
		return -1
	}

	// delete person by id 
	func DeletePersonById(personId int) int {
		db := DB()
		var person models.Person
		if db != nil {
			result := db.Where("id = ?", personId).Delete(&person)
			if result.RowsAffected > -1 {
				return int(result.RowsAffected)
			}
		}
		return -1
	}

	func UpdateMarriedStatus(personId int, marriedStatus bool) int{
		db := DB()
		var person models.Person
		if db != nil {
			result := db.Model(&person).Where("id = ?", personId).Update("is_married", marriedStatus)
			if result.RowsAffected > -1 {
				return int(result.RowsAffected)
		}
	}
		return -1
		
	}