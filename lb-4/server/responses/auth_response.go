package responses

import "github.com/dimaniko04/Go/lb-4/server/models"

type AuthResponse struct {
	User  models.User
	Token string
}
