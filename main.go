// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var addr = flag.String("addr", ":1323", "http service address")

func main() {

	flag.Parse()

	hub := newHub()
	go hub.run()

	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.Recover())

	e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "[${status}] ${error} host:${host} uri:${uri} method:${method} user_agent:${user_agent}\n",
	// }))

	e.GET("/", healthCheck)

	e.GET("/ws/:userName", func(c echo.Context) error {
		return serveWs(hub, c, c.Param("userName"))
	})

	e.Logger.Fatal(e.Start(*addr))

}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
