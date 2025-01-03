package requests

type CheckoutRequest struct {
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Card           string `json:"card"`
	CardExpiryDate string `json:"card_expiry_date"`
	CardCvv        string `json:"card_cvv"`
}
