package domain

type Passport struct {
	Message string
	Token   string
}

func BuildPassport(message, token string) Passport {
	return Passport{
		message,
		token,
	}
}
