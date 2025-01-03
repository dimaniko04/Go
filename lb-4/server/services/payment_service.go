package services

import (
	"bytes"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dimaniko04/Go/lb-4/server/requests"
)

type PaymentService interface {
	Payment(requests.PaymentRequest) error
}

type paymentService struct {
	db               *sql.DB
	liqpayPrivateKey string
}

func (s *paymentService) Payment(request requests.PaymentRequest) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	data := base64.StdEncoding.EncodeToString(jsonData)
	combinedString := s.liqpayPrivateKey + data + s.liqpayPrivateKey

	hash := sha1.New()
	hash.Write([]byte(combinedString))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return sendLiqPayRequest(data, signature)
}

func sendLiqPayRequest(data, signature string) error {
	client := &http.Client{}
	url := "https://www.liqpay.ua/api/request"
	formData := fmt.Sprintf("data=%s&signature=%s", data, signature)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(formData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("payment failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if !strings.Contains(string(body), `"status":"success"`) {
		return errors.New("payment failed")
	}

	return nil
}
