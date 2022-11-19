package model

/*
type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}*/

type Users struct {
	Id     uint64 `gorm:"not null;autoIncrement;primaryKey" json:"id"` 
	Name   string `gorm:"not null" json:"name"`
	Passwd string `gorm:"not null" json:"passwd"`
}
type Todos struct {
	Id      uint64 `gorm:"not null;autoIncrement;primaryKey"`
	User_id uint64 `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
