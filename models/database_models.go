package models

type House struct {
	Id           int		`gorm:"primaryKey"`
	Name         string
	Animal     	 string
	Motto        string
}

type Person struct {
	Id           int	 	`gorm:"primaryKey"`
	Name         string
	IsMarried    bool
	HouseId      int     	`gorm:"foreignKey:HouseId"`
}
