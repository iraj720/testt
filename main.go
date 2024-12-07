package main

import (
	"net/http"
	"test/handler"
	"test/metrics"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var Calls metrics.Metrics

func main() {
	e := echo.New()

	Calls = metrics.NewMetrics()

	helloController := handler.NewHelloHandler(Calls)
	goodbyeController := handler.NewGoodbyeHandler(Calls)

	e.GET("/hello", helloController.Hello)
	e.GET("/goodbye", goodbyeController.Goodbye)

	e.GET("/report", Report)

	if err := e.Start("127.0.0.1:8080"); err != nil {
		logrus.Fatalf("failed to start http server: %s", err.Error())
	}
}

func Report(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Calls.GetBatchReport())
}
