package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/ping [get]
func (c *Endpoint) PingExample(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

// CalcExample godoc
// @Summary calc example
// @Description plus
// @Tags example
// @Accept json
// @Produce json
// @Param val1 query int true "used for calc"
// @Param val2 query int true "used for calc"
// @Success 200 {integer} integer "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/calc [get]
func (c *Endpoint) CalcExample(ctx echo.Context) error {
	val1, err := strconv.Atoi(ctx.QueryParam("val1"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	val2, err := strconv.Atoi(ctx.QueryParam("val2"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	ans := val1 + val2
	return ctx.String(http.StatusOK, fmt.Sprintf("%d", ans))
}

// PathParamsExample godoc
// @Summary path params example
// @Description path params
// @Tags example
// @Accept json
// @Produce json
// @Param group_id path int true "Group ID"
// @Param account_id path int true "Account ID"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/groups/{group_id}/accounts/{account_id} [get]
func (c *Endpoint) PathParamsExample(ctx echo.Context) error {
	groupID, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	accountID, err := strconv.Atoi(ctx.Param("account_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return ctx.String(http.StatusOK, fmt.Sprintf("group_id=%d account_id=%d", groupID, accountID))
}

// HeaderExample godoc
// @Summary custome header example
// @Description custome header
// @Tags example
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/header [get]
func (c *Endpoint) HeaderExample(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Request().Header.Get("Authorization"))
}

// SecuritiesExample godoc
// @Summary custome header example
// @Description custome header
// @Tags example
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Security ApiKeyAuth
// @Security OAuth2Implicit[admin, write]
// @Router /examples/securities [get]
func (c *Endpoint) SecuritiesExample(ctx echo.Context) error {
	return nil
}

// AttributeExample godoc
// @Summary attribute example
// @Description attribute
// @Tags example
// @Accept json
// @Produce json
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)
// @Param default query string false "string default" default(A)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/attribute [get]
func (c *Endpoint) AttributeExample(ctx echo.Context) error {
	return ctx.String(http.StatusOK, fmt.Sprintf("enumstring=%s enumint=%s enumnumber=%s string=%s int=%s default=%s",
		ctx.QueryParam("enumstring"),
		ctx.QueryParam("enumint"),
		ctx.QueryParam("enumnumber"),
		ctx.QueryParam("string"),
		ctx.QueryParam("int"),
		ctx.QueryParam("default"),
	))
}
