package wrapper

type BankAccount struct {
	RefundState   bool   `json:"refundState"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	Payee         string `json:"payee"`
}

type RefundWrapper struct {
	Timestamp   string      `json:"timestamp"`
	ReservedID  string      `json:"reservedId"`
	BankAccount BankAccount `json:"bankAccount"`
}
