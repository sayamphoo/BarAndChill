package entity

import (
	"encoding/json"
)

const (
	EntityDrinkName = "drink"
)

type Drink struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint16 `json:"price"`
}

func GetDrinkEntity() *[]Drink {

	jsonData := `
	[
		{"_id": "1", "name": "ช้าง", "description": "เบียร์สดยอดนิยมในประเทศไทย", "price": 100},
		{"_id": "2", "name": "สิงห์", "description": "เบียร์ยอดนิยมอีกยี่ห้อหนึ่งในไทย", "price": 120},
		{"_id": "3", "name": "ลีโอ", "description": "เบียร์สัญชาติไทยอีกยี่ห้อหนึ่ง", "price": 110},
		{"_id": "4", "name": "ไฮเนเก้น", "description": "เบียร์ระดับโลกที่ได้รับความนิยมในประเทศไทย", "price": 150},
		{"_id": "5", "name": "ช้าง ช้างน้อย", "description": "เบียร์สดสำหรับเด็กอายุ 15 ปีขึ้นไป", "price": 70},
		{"_id": "6", "name": "ลีโอ ไวท์", "description": "เบียร์ขาวยอดนิยมในประเทศไทย", "price": 120},
		{"_id": "7", "name": "สิงห์ ค็อกเทล", "description": "เบียร์ผสมค็อกเทล", "price": 130},
		{"_id": "8", "name": "เบียร์ช้าง โซดา", "description": "เบียร์ผสมโซดา", "price": 100},
		{"_id": "9", "name": "เบียร์สิงห์ โซดา", "description": "เบียร์ผสมโซดา", "price": 110},
		{"_id": "10", "name": "เบียร์ลีโอ โซดา", "description": "เบียร์ผสมโซดา", "price": 120},
		{"_id": "11", "name": "แสงโสม", "description": "เหล้าขาวยอดนิยมในประเทศไทย", "price": 150},
		{"_id": "12", "name": "เมจิก", "description": "เหล้าขาวอีกยี่ห้อหนึ่งที่นิยมในประเทศไทย", "price": 130},
		{"_id": "13", "name": "สาเก ฮอกไกโด อิชิบาชิ", "description": "สาเกยอดนิยมจากญี่ปุ่น", "price": 500},
		{"_id": "14", "name": "วิสกี้ จอห์นนี่ วอล์กเกอร์", "description": "วิสกี้ระดับโลกที่ได้รับความนิยมในประเทศไทย", "price": 1,000},
		{"_id": "15", "name": "บรั่นดี มาร์ตินี่", "description": "บรั่นดีระดับโลกที่ได้รับความนิยมในประเทศไทย", "price": 1,500}
	]
	`

	var entity []Drink

	err := json.Unmarshal([]byte(jsonData), &entity)
	if err != nil {
		return nil
	}

	return &entity
}
