package billing

type CreatePayment struct {
	BillingId string `json:"billingId"`
}
type CreatePaymentResponse struct {
	Username string `json:"username"`
	CreatePayment
}
