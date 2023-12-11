package wrapper

type TableReservationWrapper struct {
	UserID    string `json:"user_id"`
	DrinkID   string `json:"drink_id"`
	TableID   string `json:"table_id"`
	Arrival   string `json:"arrival"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"` // คอลัมน์ timestamp ใช้เก็บข้อมูลเวลา
	Statement string `json:"statement"` // คอลัมน์ statement ใช้เก็บข้อมูลข้อความ
}
