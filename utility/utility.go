package utility

import (
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/security/jwt"
	"time"
)

func GetTimeNow() string {
	return (time.Now()).Format("2006-01-02 15:04:05.999999999")
}

func BuildPassport(userID *string) (*domain.Passport, error) {
	pass, err := jwt.BuildJwt(userID)
	if err != nil {
		return nil, err
	}

	return &domain.Passport{
		Message: "SUCCESS",
		Token:   *pass,
	}, nil
}
