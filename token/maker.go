package token

import (
	"time"
)

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	// Gotta create a payload struct for verify token
	VerifyToken(token string) (*Payload, error)
}
