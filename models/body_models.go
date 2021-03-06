package models

type InsertHouseRequestBody struct{
	Name         string		`json:"name" binding:"required"`
	Animal     	 string		`json:"animal" binding:"required"`
	Motto        string		`json:"motto" binding:"required"`
}

type InsertPersonRequestBody struct{
	Name         string     `json:"name" binding:"required"`
	HouseId      int     	`json:"houseId" binding:"required"`
	IsMarried    bool		`json:"isMarried"`
}