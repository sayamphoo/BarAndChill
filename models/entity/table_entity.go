package entity

const (
	EntityTableName = "table"
)

type TableEntity struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Chair uint8  `json:"chair"`
}

func GetTable() []TableEntity {
	tableEntities := []TableEntity{
		{ID: "1", Name: "Table 1", Chair: 4},
		{ID: "2", Name: "Table 2", Chair: 6},
		{ID: "3", Name: "Table 3", Chair: 2},
		{ID: "4", Name: "Table 4", Chair: 8},
		{ID: "5", Name: "Table 5", Chair: 5},
		{ID: "6", Name: "Table 6", Chair: 4},
		{ID: "7", Name: "Table 7", Chair: 6},
		{ID: "8", Name: "Table 8", Chair: 10},
		{ID: "9", Name: "Table 9", Chair: 3},
		{ID: "10", Name: "Table 10", Chair: 7},
	}

	return tableEntities
}
