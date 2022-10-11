package billing

import (
	"net/http"
	"tdg-advanced-go-day-1/workshop/app/base"

	"github.com/labstack/echo/v4"
)

func GetBillingHandler(service *BillingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Param("username")
		username := service.GetBillingByUsername(u)
		response := base.NewResponse(username, "success")
		return c.JSON(http.StatusOK, response)
	}
}

func CreateBillingHandler(service *BillingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		payload := new(CreatePayment)
		if err := c.Bind(payload); err != nil {
			return c.JSON(http.StatusBadRequest, base.NewResponse(nil, "bad request").WithError(err))
		}
		res := service.CreateBilling(id, payload)

		return c.JSON(http.StatusOK, base.NewResponse(res, "created"))
	}
}
