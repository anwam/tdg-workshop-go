package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreatePayment struct {
	BillingId string `json:"billingId"`
}

type CreatePaymentResponse struct {
	Username string `json:"username"`
	CreatePayment
}

func main() {
	e := echo.New()

	e.GET("/billing/:username", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", username))
	})

	e.POST("/billing/:username", func(c echo.Context) error {
		username := c.Param("username")
		payload := new(CreatePayment)
		if err := c.Bind(payload); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		response := CreatePaymentResponse{
			Username:      username,
			CreatePayment: *payload,
		}
		return c.JSON(http.StatusCreated, response)
	})
	e.Logger.Fatal(e.Start(":3030"))
}
