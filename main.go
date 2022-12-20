// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package main

import (
	"flag"

	api "demo-echo-v4/api"
	_ "demo-echo-v4/docs"
	websocket "demo-echo-v4/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	swagger "github.com/swaggo/echo-swagger"
)

var addr = flag.String("addr", ":1323", "http service address")

func main() {

	flag.Parse()

	hub := websocket.NewHub()
	go hub.Run()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", api.HealthCheck)
	api.HandlerMapping(e)

	// Swagger UI
	e.GET("/swagger/*", swagger.WrapHandler)

	// Websocket
	e.GET("/ws/:userName", func(c echo.Context) error {
		return websocket.ServeWs(hub, c)
	})

	e.Logger.Fatal(e.Start(*addr))

}
