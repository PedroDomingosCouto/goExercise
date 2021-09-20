package helpers

import "goExercise/models"

// Query to get houses from the database
func GetAllHouses() *[]models.House {
	var houses []models.House

	db := DB()

	if db != nil {
		db.Find(&houses)
	}
	return &houses
  }

	// Query to get House by id
	func GetHouseById(houseId int) *models.House {
		var house models.House
		db := DB()

		if db != nil {
			result := db.Where("id = ?", houseId ).First(&house)
			if result.RowsAffected > 0{
				return &house
			}
			
		}
		return nil
  	}

  // insert a house into database
  //returns the inserted row id or -1 if operation failed
  func CreateHouse(houseBody models.InsertHouseRequestBody) *models.House {
	db := DB()

	parsedHouse:=CreateHouseObject(houseBody)

	result := db.Create(&parsedHouse)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &parsedHouse

  }

  // update house by id
  func UpdateOrCreateHouse(houseId int, houseRequest *models.InsertHouseRequestBody) *models.House{
	db := DB()

	house:=CreateHouseObject(*houseRequest)
	if db != nil {
		resultUpdate := db.Model(&house).
		Where("id = ? ",houseId).
		Updates(map[string]interface{}{"name": house.Name, "animal": house.Animal, "motto": house.Motto});
		
		if resultUpdate != nil {
			if resultUpdate.RowsAffected == 0 {
				resultCreate := db.Create(&house)
				if resultCreate.RowsAffected > 0  {
					return GetHouseById(house.Id)
				}
				return nil
			} else if resultUpdate.RowsAffected == 1{
				return GetHouseById(houseId)
			}
		}
	}
	return nil
}


  // delete house by id 
  // validar qtd de membros com aquela casa
func DeleteHouseById(houseId int) int {
	db := DB()
	var house models.House
	if db != nil {
		result := db.Where("id = ?",houseId).Delete(&house)
		if result.RowsAffected == 1 {
			return 1
		}
	}
	return -1
}


//Get a list with persons in house
func GetQuantityPersonByHouseID(houseId int) int{
	db := DB()
	var persons []models.Person

	if db != nil {
		result:=db.Where("house_id = ?", houseId).Find(&persons)
		return int(result.RowsAffected)
	}
	
	return -1
}