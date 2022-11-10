package app

import (
	"flybitch/utils"
	//"net/http"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var e *echo.Echo

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes()
	e.Validator = &utils.CustomValidator{}

	logrus.Info("echo framework initialized")
}




func StartServer() {
    e.Logger.Fatal(e.Start("127.0.0.1:80"))  // 启动服务，注意默认端口80不能省略

}
