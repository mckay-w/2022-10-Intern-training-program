package model

/*
type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}*/

type Users struct {
	Id     uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Name   string `gorm:"not null"`
	Passwd string `gorm:"not null"`
}
type Todos struct {
	Id      uint64 `gorm:"not null;autoIncrement;primaryKey"`
	User_id uint64 `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
