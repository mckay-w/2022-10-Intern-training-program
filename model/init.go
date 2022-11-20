package model

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
	tet := viper.Get("username")
	username := viper.GetString("username")
	dbArgs := username + ":" + viper.GetString("password") +
		"@(localhost)/" + viper.GetString("db_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(tet)
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}

type request struct {
	keyword string `josn:"keyword" `
	value   string `josn:"value"`
}

// 键值对查找
func Myfind(c echo.Context) (err error) {
	var result Users
	rqst := new(request)
	if err = c.Bind(rqst); err != nil {
		return
	}
	DB.Where(rqst.keyword+" = ?", rqst.value).Find(&result)
	return c.JSON(http.StatusOK, result)
}

type Adusers struct {
	Name   string `gorm:"not null" json:"name" form:"name" query:"name"`
	Passwd string `gorm:"not null" json:"passwd" form:"passwd" query:"passwd"`
}

// add user to sql
func Adduser(c echo.Context) (err error) {
	//var newuser Users;
	newuser := new(Users)
	if err = c.Bind(newuser); err != nil {
		return
	}
	//str:="\"Keyword\":\""+s.Keyword+"\",\"Value\":\""+info.Value+"\""
	// 接口参数校验，防攻击
	validate := validator.New()
	if err := validate.Struct(newuser); err != nil {
		fmt.Println("Args not allow.", err)
		logrus.Panic("Args not allow.", err)
	}
	DB.Create(&newuser)
	return c.JSON(http.StatusOK, newuser)
}

// delete user from sql
func Deteleuser(c echo.Context) (err error) {
	duser := new(Users)
	if err = c.Bind(duser); err != nil {
		return
	}
	//str:="\"Keyword\":\""+.Keyword+"\",\"Value\":\""+info.Value+"\""
	// 接口参数校验，防攻击
	validate := validator.New()
	if err := validate.Struct(duser); err != nil {
		fmt.Println("Args not allow.", err)
		logrus.Panic("Args not allow.", err)
	}
	DB.Delete(&duser)
	return c.JSON(http.StatusOK, duser)
}

// modify user field in sql
func Modify(user Users, modi request) {
	if modi.keyword == "id" {
		//user.id=uint(modi.value)
	} else if modi.keyword == "name" {
		user.Name = modi.value
	} else if modi.keyword == "passwd" {
		user.Passwd = modi.value
	}

}

func getbody(c echo.Context) []byte {
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(err)
		logrus.Panic(err)
	}
	return body
}


