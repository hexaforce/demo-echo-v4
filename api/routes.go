package api

import (
	"net/http"

	endpoint "demo-echo-v4/api/endpoint"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func HandlerMapping(e *echo.Echo) {

	handler := endpoint.NewEndpoint()

	v1 := e.Group("/v1")
	accounts := v1.Group("/accounts")
	{
		{
			accounts.GET(":id", handler.ShowAccount)
			accounts.GET("", handler.ListAccounts)
			accounts.POST("", handler.AddAccount)
			accounts.DELETE(":id", handler.DeleteAccount)
			accounts.PATCH(":id", handler.UpdateAccount)
			accounts.POST(":id/images", handler.UploadAccountImage)
		}
		bottles := v1.Group("/bottles")
		{
			bottles.GET(":id", handler.ShowBottle)
			bottles.GET("", handler.ListBottles)
		}
		admin := v1.Group("/admin")
		{
			admin.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					if len(c.Request().Header.Get("Authorization")) == 0 {
						return echo.NewHTTPError(http.StatusUnauthorized, "Authorization is required Header")
					}
					return nil
				}
			})
			admin.POST("/auth", handler.Auth)
		}
		examples := v1.Group("/examples")
		{
			examples.GET("ping", handler.PingExample)
			examples.GET("calc", handler.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", handler.PathParamsExample)
			examples.GET("header", handler.HeaderExample)
			examples.GET("securities", handler.SecuritiesExample)
			examples.GET("attribute", handler.AttributeExample)
		}
	}
}
