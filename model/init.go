package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func Myfind(c echo.Context) Users {
	var result Users
	var body []byte
	body=getbody(c)
    var rqst request;
    if err := json.Unmarshal(body, &rqst); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        logrus.Panic("json unmarshal error: ", err)
    }
    DB.Where(rqst.keyword+" = ?",rqst.value).Find(&result)
    return result
}

// add user to sql
func Adduser(c echo.Context) {
	var newuser Users;
    var body []byte
    body=getbody(c)
    if err := json.Unmarshal(body, &newuser); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        logrus.Panic("json unmarshal error: ", err)
    }
    //str:="\"Keyword\":\""+.Keyword+"\",\"Value\":\""+info.Value+"\""
    // 接口参数校验，防攻击
    validate := validator.New()
    if err := validate.Struct(newuser); err != nil {
        fmt.Println("Args not allow.", err)
        logrus.Panic("Args not allow.", err)
    }
	DB.Create(&newuser)
}

// delete user from sql
func Deteleuser(c echo.Context){
	var duser Users;
    var body []byte
    body=getbody(c)
    if err := json.Unmarshal(body, &duser); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        logrus.Panic("json unmarshal error: ", err)
    }
    //str:="\"Keyword\":\""+.Keyword+"\",\"Value\":\""+info.Value+"\""
    // 接口参数校验，防攻击
    validate := validator.New()
    if err := validate.Struct(duser); err != nil {
        fmt.Println("Args not allow.", err)
        logrus.Panic("Args not allow.", err)
    }
	DB.Delete(&duser)
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

func getbody(c echo.Context) []byte{
	defer c.Request().Body.Close()
    body, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        fmt.Println(err)
        logrus.Panic(err)
    }
	return  body
}
