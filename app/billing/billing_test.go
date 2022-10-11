// go:build integration

package billing_test

import (
	"net/http"
	"net/http/httptest"
	"tdg-advanced-go-day-1/workshop/app/billing"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_SuccessWithGet(t *testing.T) {
	// Setup http server engine
	e := echo.New()
	billingServ := billing.NewBillingService()
	req := httptest.NewRequest(http.MethodGet, "/billing/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("username")
	c.SetParamValues("asdfasdf")
	if assert.NoError(t, billing.GetBillingHandler(billingServ)(c)) {
		// Assertions
		assert.NotEmpty(t, rec.Body.String())
		assert.JSONEq(t, `{"data": "asdfasdf", "error": null, "message": "success"}`, rec.Body.String())
	}
}
