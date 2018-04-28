package models

type Goods struct {
	Id int
	Name string
	Desc string
	Price float64 `orm:"digits(12);decimals(2)"`
	Inventory int
	Img string
	UpdateId int
	CreateId int
	CreateTime int64
	UpdateTime int64
	Status int8

}

