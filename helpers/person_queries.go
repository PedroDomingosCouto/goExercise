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
		result := db.Where("id = ?", personId).Find(&person)
		if result.RowsAffected > 0{
			return &person
		}
	}
	return nil
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
	func AddNewPerson(personBody models.InsertPersonRequestBody) *models.Person {
		db := DB()

		parsedPerson := CreatePersonObject(personBody)
		
		if db != nil {
			result := db.Create(&parsedPerson)
			if result.RowsAffected == 0 || result.Error != nil {
				return nil
			}
			return &parsedPerson
		}
		return nil
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

	//update maried status
	func UpdateMarriedStatus(personId int, marriedStatus bool) *models.Person{
		db := DB()
		var person models.Person
		if db != nil {
			result := db.Model(&person).Where("id = ?", personId).Update("is_married", marriedStatus)
			if result.RowsAffected > 0 {
				return GetPersonById(personId)
			}
		}
		return nil
	}

	//update house id
	func UpdateHouseId(personId int, newHouseId int) *models.Person {
		db := DB()
		var person models.Person
		if db != nil {
			result := db.Model(&person).Where("id = ?", personId).Update("house_id", newHouseId)
			if result.RowsAffected > 0 {
				return GetPersonById(personId)
			}
		}
		return nil
	}
