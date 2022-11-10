package controller

import (
	"flybitch/app/response"
	"net/http"
	"github.com/labstack/echo/v4"
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
	"github.com/go-playground/validator"
    "github.com/sirupsen/logrus"
)

type request struct{
	//keyword string `josn:"keyword"`
	Keyword string `josn:"keyword"`
	Value string `josn:"value"`
}


func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}


// return itself
func Query(c echo.Context) error {
    qry:=c.Request().URL.Query().Encode()
    if qry ==""{
        fmt.Println("No query")
        logrus.Panic("No query")
        return c.String(http.StatusBadRequest, "Bad request")
    }else {
        fmt.Println(qry)
        return response.SendResponse(c, http.StatusOK, "query", c)
    }
}

/*
get a struct from request
print it in form of josn
*/
func Analysis(c echo.Context) error{
	defer c.Request().Body.Close()
    body, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        fmt.Println(err)
        logrus.Panic(err)
    }

    var info request
    if err = json.Unmarshal(body, &info); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        logrus.Panic("json unmarshal error: ", err)
        return c.String(http.StatusInternalServerError, "Json unmarshal error")
    }
    str:="\"Keyword\":\""+info.Keyword+"\",\"Value\":\""+info.Value+"\""
    // 接口参数校验，防攻击
    validate := validator.New()
    if err = validate.Struct(info); err != nil {
        fmt.Println("Args not allow.", err)
        logrus.Panic("Args not allow.", err)
        return c.String(http.StatusBadRequest, "Bad request")
    }
    return response.SendResponse(c,http.StatusOK, "body", str)
}






