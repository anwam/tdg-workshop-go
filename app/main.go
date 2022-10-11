package main

import (
	"os"
	"tdg-advanced-go-day-1/workshop/app/billing"

	"github.com/labstack/echo/v4"
)

// var cred = flag.String("cred", "asdf1324", "Credentials")

func main() {
	// flag.Parse()
	port := os.Getenv("PORT")
	e := echo.New()

	billingServ := billing.NewBillingService()

	e.GET("/billing/:username", billing.GetBillingHandler(billingServ))
	e.POST("/billing/:id", billing.CreateBillingHandler(billingServ))

	e.Logger.Fatal(e.Start(":" + port))
}
