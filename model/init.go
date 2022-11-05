package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	err := DB.AutoMigrate(&Foo{}) // TODO: add table structs here
	if err != nil {
		logrus.Fatal(err)
	}
	//var tablename1="users_go";
	//var tablename2="todo_go";
	// newly build two tables
	DB.AutoMigrate(&Users{})
	DB.AutoMigrate(&Todos{})


}

func connectDatabase() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	//if err := viper.ReadInConfig(); err != nil {
	//logrus.Panic(err)
	//}

	//loginInfo := viper.GetStringMapString("sql")

	/*dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
	"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"*/
	dbArgs := "todo" + ":" + "todo123" + "@(localhost)/" + "todo" + "?charset=utf8mb4&parseTime=True&loc=Local"
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
func myfind(rqst request) Users {
	var result Users;
    DB.Where(rqst.keyword+" = ?",rqst.value).Find(&result)
    return result
}


func adduser(newuser Users) {
	DB.Create(&newuser)
}

func deleteuser(user Users){
	DB.Delete(&user)
}

func dodify(user Users, modi request)
{
	user.(modi.keyword)=modi.value
}