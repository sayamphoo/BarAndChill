package entity

const EntityBank = "bankrefund"

type BankAccount struct {
	RefundState   bool   `json:"refundState"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	Payee         string `json:"payee"`
}

type RefundEntity struct {
	ID          string      `json:"_id"`
	Timestamp   string      `json:"timestamp"`
	ReservedID  string      `json:"reservedId"`
	BankAccount BankAccount `json:"bankAccount"`
}
