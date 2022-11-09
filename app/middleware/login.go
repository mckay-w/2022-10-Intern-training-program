package middleware

import (
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", 114514)
		return next(c)
	}
}


/*
func startEchoServer() error {
    e := echo.New()

    // 路由表
    e.GET("/hello", handleGet)
    e.POST("/world", handlePost)

    // 启动服务
    if err := e.Start("127.0.0.1:8080"); err != nil {
        fmt.Println("Echo start error: ", err)
        return err
    }

    return nil
}

func main() {
    if err := startEchoServer(); err != nil {
        fmt.Println("Start Server error.", err)
        return
    }
}
*/
// middleware
//func Mw(next echo.HandlerFunc) echo.HandlerFunc {
	
//}
