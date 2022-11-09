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
	//"flybitch/app"
)


// post 请求json字符串格式定义
type PostInfo struct {
    Code   int    `json:"Code" validate:"required,min=0,max=10"`  // 必须携带，而且是(0,10]的整数
    Msg    string `json:"Msg" validate:"required,excludesall=!@#$%^&*()_-{}"` // 防特殊字符输入造成攻击风险
}

type request struct{
	//keyword string `josn:"keyword"`
	Keyword string `josn:"keyword"`
	Value string `josn:"value"`
}


func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}
/*
// get请求处理
func handleGet(c echo.Context) error {
    return c.String(http.StatusOK, "Hello")
}

// post请求处理
func handlePost(c echo.Context) error {
    defer c.Request().Body.Close()
    body, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        fmt.Println(err)
    }

    var info PostInfo
    if err = json.Unmarshal(body, &info); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        return c.String(http.StatusInternalServerError, "Json unmarshal error")
    }

    // 接口参数校验，防攻击
    validate := validator.New()
    if err = validate.Struct(info); err != nil {
        fmt.Println("Args not allow.", err)
        return c.String(http.StatusBadRequest, "Bad request")
    }

    return c.String(http.StatusOK, "Post json data ok.")
}*/

type CustomContext struct {
	echo.Context
}


// return itself
func Query(c echo.Context) error {
    qry:=c.Request().URL.Query().Encode()
    if qry ==""{
        fmt.Println("No query")
        return c.String(http.StatusBadRequest, "Bad request")
    }else {
        fmt.Println(qry)
        return response.SendResponse(c, http.StatusOK, "query", c)
    }
}


func Analysis(c echo.Context) error{
	defer c.Request().Body.Close()
    body, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        fmt.Println(err)
    }

    var info request
    if err = json.Unmarshal(body, &info); err != nil {  // 将字节数组转换成struct类型
        fmt.Println("json unmarshal error: ", err)
        return c.String(http.StatusInternalServerError, "Json unmarshal error")
    }
	fmt.Println(info)

    // 接口参数校验，防攻击
    validate := validator.New()
    if err = validate.Struct(info); err != nil {
        fmt.Println("Args not allow.", err)
        return c.String(http.StatusBadRequest, "Bad request")
    }


	rqst:=new(request)
	if err := c.Bind(rqst); err != nil {
		return err
	}
	jrqst, _ := json.Marshal(rqst)
	var out bytes.Buffer
	json.Indent(&out, jrqst,"", "\t")
	fmt.Println(jrqst)
	return c.JSON(http.StatusOK, rqst)
}




type CustomBinder struct {}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// 你也许会用到默认的绑定器
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	// 做你自己的实现

	return
}
/*
func run(){
	resp, err := http.Get("http://httpbin.org/get")
	defer resp.Body.Close()
    // 200 OK
    fmt.Println(resp.Status)
    fmt.Println(resp.Header)

    buf := make([]byte, 1024)
    for {
        // 接收服务端信息
        n, err := resp.Body.Read(buf)
        if err != nil && err != io.EOF {
            fmt.Println(err)
            return
        } else {
            res := string(buf[:n])
            fmt.Println(res)
            break
        }
    }
	//http.HandleFunc("/go", myHandler);
	

}*/


