package handler

import (
	"net/http"
	"strconv"

	schemas "demo-echo-v4/api/schemas"

	"github.com/labstack/echo/v4"
)

// ShowBottle godoc
// @Summary Show a bottle
// @Description get string by ID
// @ID get-string-by-int
// @Tags bottles
// @Accept  json
// @Produce  json
// @Param  id path int true "Bottle ID"
// @Success 200 {object} schemas.Bottle
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles/{id} [get]
func (c *Endpoint) ShowBottle(ctx echo.Context) error {
	id := ctx.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	bottle, err := schemas.BottleOne(bid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, bottle)
}

// ListBottles godoc
// @Summary List bottles
// @Description get bottles
// @Tags bottles
// @Accept  json
// @Produce  json
// @Success 200 {array} schemas.Bottle
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles [get]
func (c *Endpoint) ListBottles(ctx echo.Context) error {
	bottles, err := schemas.BottlesAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, bottles)
}
