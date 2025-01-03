package requests

type PaymentRequest struct {
	PublicKey    string `json:"public_key"`
	Version      string `json:"version"`
	Action       string `json:"action"`
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	Description  string `json:"description"`
	Card         string `json:"card"`
	CardExpYear  string `json:"card_exp_year"`
	CardExpMonth string `json:"card_exp_month"`
	CardCvv      string `json:"card_cvv"`
}
