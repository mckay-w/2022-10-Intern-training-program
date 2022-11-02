package model

type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}


type users struct{
	id     uint primary key
	name   string
	passwd string
}

type Todos struct{
	id       uint primary key
	user_id  uint
	title    string
	content  string
}