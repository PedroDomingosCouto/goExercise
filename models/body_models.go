package models

type InsertHouseRequestBody struct{
	Name         string		`json:"name" binding:"required"`
	Animal     	 string		`json:"animal" binding:"required"`
	Motto        string		`json:"motto" binding:"required"`
}