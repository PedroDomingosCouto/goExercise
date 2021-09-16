package models

import (
	"fmt"
	"os"
)

type House struct {
	Id           int		`gorm:"primaryKey"`
	Name         string
	Animal     	 string
	Motto        string
}

func (House) TableName() string {
	return fmt.Sprintf("%s.house", os.Getenv("DB_SCHEMA"))
}

type Person struct {
	Id           int	 	`gorm:"primaryKey"`
	Name         string
	IsMarried    bool
	HouseId      int     	`gorm:"foreignKey:HouseId"`
}

func (Person) TableName() string {
	return fmt.Sprintf("%s.person", os.Getenv("DB_SCHEMA"))
}