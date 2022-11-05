package controller

import (
	"flybitch/app/response"
	"flybitch/model"
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo/v4"
	//"flybitch/app"
)

func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}



// return itself
func query(qry string) string {
	return qry
}

type request struct{
	keyword string
	value string
}



// analyse 
func analyse(rqst request) error {



	
	myfind()
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

func run()
{
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
	http.HandleFunc("/go", myHandler);
	

}