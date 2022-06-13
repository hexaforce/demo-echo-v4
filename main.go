// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package main

import (
	"flag"

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

	e.GET("/", healthCheck)

	e.GET("/ws/:userName", func(c echo.Context) error {
		return serveWs(hub, c)
	})

	e.Logger.Fatal(e.Start(*addr))

}
