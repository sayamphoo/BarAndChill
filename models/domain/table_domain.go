package domain

type TableDomain struct {
	Id    string `json:"_id"`
	Name  string `json:"name"`
	Chair uint8  `json:"chair"`
	State bool
}
