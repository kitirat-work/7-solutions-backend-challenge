package main

import (
	"piefiredire/baconipsum"
	"piefiredire/handler"
	"piefiredire/service"

	"github.com/imroc/req/v3"
	"github.com/labstack/echo/v4"
)

func main() {
	client := req.C()
	baconipsumService := baconipsum.NewBaconipsum(client)
	srv := service.NewService(baconipsumService)
	h := handler.NewHandler(srv)

	e := echo.New()
	e.GET("/beef/summary", h.BeefSummary)

	e.Logger.Fatal(e.Start(":8080"))
}
