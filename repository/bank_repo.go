package repository

import (
	"encoding/json"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository/repo"
)

type BankRepo struct {
}

var repoBank = repo.Database{
	Index: entity.EntityBank,
}

func (m *BankRepo) Sava(wp *wrapper.RefundWrapper) (string, error) {
	return repoBank.RepoSave(wp)
}

func (m *BankRepo) FindByIdReservation(id string) *entity.RefundEntity {
	data, _ := repoBank.RepoFindByWord("reservedId", id)

	var entity entity.RefundEntity
	if err := json.Unmarshal(data.Hits[0].Source_, &entity); err != nil {
		return nil
	}

	return &entity
}
