package entity

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	EntityReserveTableName = "reserve_table"
)

type ReserveTable struct {
	ID        string `json:"_id"`      // คอลัมน์ id ใช้เป็น primary key
	UserID    string `json:"user_id"`  // คอลัมน์ user_id ใช้เป็น foreign key
	DrinkID   string `json:"drink_id"` // คอลัมน์ promotion_id ใช้เป็น foreign key
	TableID   string `json:"table_id"` // คอลัมน์ table_id ใช้เป็น foreign key
	Arrival   string `json:"arrival"`
	Timestamp string `json:"timestamp"` // คอลัมน์ timestamp ใช้เก็บข้อมูลเวลา
	Statement string `json:"statement"` // คอลัมน์ statement ใช้เก็บข้อมูลข้อความ
	Status    string `json:"status"`    //รอยืนยันจากเจ้าของร้าน
}

func HitsToReserveTable(wr types.HitsMetadata) (*[]ReserveTable, error) {
	entityList := []ReserveTable{}
	rawData := wr.Hits

	for _, data := range rawData {
		var e ReserveTable
		if err := json.Unmarshal(data.Source_, &e); err != nil {
			continue
		}
		e.ID = data.Id_
		entityList = append(entityList, e)
	}

	return &entityList, nil
}
