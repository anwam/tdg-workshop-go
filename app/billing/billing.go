package billing

type BillingService struct {
}

func NewBillingService() *BillingService {
	return &BillingService{}
}

func (BillingService) GetBillingByUsername(username string) string {
	return username
}

func (BillingService) CreateBilling(id string, payload *CreatePayment) CreatePaymentResponse {
	response := CreatePaymentResponse{
		Username:      id,
		CreatePayment: *payload,
	}
	return response
}
