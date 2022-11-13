package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	DB.AutoMigrate(&Users{})
	DB.AutoMigrate(&Todos{})
}

func connectDatabase() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
	logrus.Panic(err)
	}
	tet:=viper.Get("username")
	username:=viper.GetString("username")
	dbArgs:=username + ":" + viper.GetString("password")+
	"@(localhost)/" + viper.GetString("db_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(tet)
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}

type request struct{
	keyword string
	value string
}

// 键值对查找
func Myfind(rqst request) Users {
	var result Users;
    DB.Where(rqst.keyword+" = ?",rqst.value).Find(&result)
    return result
}

// add user to sql
func Adduser(newuser Users) {
	DB.Create(&newuser)
}

// delete user from sql
func Deteleuser(user Users){
	DB.Delete(&user)
}

// modify user field in sql
func Modify(user Users, modi request) {

	if(modi.keyword=="id"){
		//user.id=uint(modi.value)
	}else if(modi.keyword=="name"){
		user.Name=modi.value
	}else if(modi.keyword=="passwd"){
		user.Passwd=modi.value
	}
	
}