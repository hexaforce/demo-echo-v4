package handler

import (
	"fmt"
	"net/http"
	"strconv"

	schemas "demo-echo-v4/api/schemas"

	"github.com/labstack/echo/v4"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} schemas.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
func (c *Endpoint) ShowAccount(ctx echo.Context) error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	account, err := schemas.AccountOne(aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(email)
// @Success 200 {array} schemas.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
func (c *Endpoint) ListAccounts(ctx echo.Context) error {
	q := ctx.QueryParam("q")
	accounts, err := schemas.AccountsAll(q)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
// @Summary Add a account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body schemas.AddAccount true "Add account"
// @Success 200 {object} schemas.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func (c *Endpoint) AddAccount(ctx echo.Context) error {
	var addAccount schemas.AddAccount
	if err := ctx.Bind(&addAccount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	if err := addAccount.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	account := schemas.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	account.ID = lastID
	return ctx.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
// @Summary Update a account
// @Description Update by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  account body schemas.UpdateAccount true "Update account"
// @Success 200 {object} schemas.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [patch]
func (c *Endpoint) UpdateAccount(ctx echo.Context) error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	var updateAccount schemas.UpdateAccount
	if err := ctx.Bind(&updateAccount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	account := schemas.Account{
		ID:   aid,
		Name: updateAccount.Name,
	}
	err = account.Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
// @Summary Update a account
// @Description Delete by account ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID" Format(int64)
// @Success 204 {object} schemas.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [delete]
func (c *Endpoint) DeleteAccount(ctx echo.Context) error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	err = schemas.Delete(aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}
	return ctx.JSON(http.StatusNoContent, nil)
}

// UploadAccountImage godoc
// @Summary Upload account image
// @Description Upload file
// @Tags accounts
// @Accept  multipart/form-data
// @Produce  json
// @Param  id path int true "Account ID"
// @Param file formData file true "account image"
// @Success 200 {object} controller.Message
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id}/images [post]
func (c *Endpoint) UploadAccountImage(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload compleate userID=%d finename=%s", id, file.Filename)})
}
