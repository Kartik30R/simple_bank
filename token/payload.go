package token

import (
 	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// GetExpirationTime implements jwt.Claims.
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetIssuedAt implements jwt.Claims.
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetIssuer implements jwt.Claims.
func (payload *Payload) GetIssuer() (string, error) {
	panic("unimplemented")
}

// GetNotBefore implements jwt.Claims.
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	panic("unimplemented")
}

// GetSubject implements jwt.Claims.
func (payload *Payload) GetSubject() (string, error) {
	panic("unimplemented")
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {

	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// GetAudience implements jwt.Claims.
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	panic("unimplemented")
}

