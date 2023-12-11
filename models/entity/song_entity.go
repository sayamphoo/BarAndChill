package entity

import (
	"encoding/json"
	"sort"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	EntitySongRequestEntityName = "song_request"
)

type SongRequestEntity struct {
	ID        string `json:"_id"`       // คอลัมน์ id ใช้เป็น primary key
	UserID    string `json:"user_id"`   // คอลัมน์ user_id ใช้เป็น foreign key
	NameSong  string `json:"name_song"` // คอลัมน์ name_song ใช้เก็บชื่อเพลง
	NameUser  string `json:"nameUser"`
	URL       string `json:"url"`       // คอลัมน์ url ใช้เก็บ URL ของเพลง
	Message   string `json:"message"`   // คอลัมน์ message ใช้เก็บข้อความ
	Timestamp string `json:"timestamp"` // คอลัมน์ timestamp ใช้เก็บข้อมูลเวลา
	State     string `json:"state"`     // คอลัมน์ state ใช้เก็บสถานะ
}

func HitsToEntitySong(wr types.HitsMetadata) *[]SongRequestEntity {
	entityList := []SongRequestEntity{}
	rawData := wr.Hits

	for _, data := range rawData {
		var e SongRequestEntity
		if err := json.Unmarshal(data.Source_, &e); err != nil {
			continue
		}
		e.ID = data.Id_
		entityList = append(entityList, e)
	}

	sort.Slice(entityList, func(i, j int) bool {
		return entityList[i].Timestamp < entityList[j].Timestamp
	})

	return &entityList
}
