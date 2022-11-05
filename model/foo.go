package model

type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}


type Users struct{
	id     uint 
	name   string
	passwd string
}


type Todos struct{
	id       uint 
	user_id  uint
	title    string
	content  string
}