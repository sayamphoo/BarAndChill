package entity

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	EntityMemberName = "member"
)

type MemberEntity struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone" "`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func HitsToMemberEntity(wr types.HitsMetadata) (*MemberEntity, []MemberEntity, error) {

	entityList := []MemberEntity{}
	rawData := wr.Hits

	for _, data := range rawData {
		var e MemberEntity
		if err := json.Unmarshal(data.Source_, &e); err != nil {
			continue
		}
		e.Id = data.Id_
		entityList = append(entityList, e)
	}

	return &entityList[0], entityList, nil
}
