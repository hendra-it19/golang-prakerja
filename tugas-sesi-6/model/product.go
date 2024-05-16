package model

type Product struct {
	Id    uint   `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(255)"`
	Price uint   `gorm:"type:int"`
}
