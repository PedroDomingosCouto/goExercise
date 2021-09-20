package models

type HousePersons struct{
	House  		
	Persons []Person
}

type PersonHouse struct{
	Person 
	House House
}

