package handler

import (
	"net/http"
	"test/metrics"

	"github.com/labstack/echo"
)

type HelloHandler struct {
	calls_metrics metrics.Metrics
}

func NewHelloHandler(calls_metrics metrics.Metrics) HelloHandler {
	return HelloHandler{
		calls_metrics: calls_metrics,
	}
}

func (g *HelloHandler) Hello(ctx echo.Context) error {
	g.calls_metrics.IncCalls(ctx.Path())
	return ctx.String(http.StatusOK, "hello")
}

type GoodbyeHandler struct {
	calls_metrics metrics.Metrics
}

func NewGoodbyeHandler(calls_metrics metrics.Metrics) GoodbyeHandler {
	return GoodbyeHandler{
		calls_metrics: calls_metrics,
	}
}

func (g *GoodbyeHandler) Goodbye(ctx echo.Context) error {
	g.calls_metrics.IncCalls(ctx.Path())
	return ctx.String(http.StatusOK, "goodbye")
}
