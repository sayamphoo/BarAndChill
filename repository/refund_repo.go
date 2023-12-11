package repository

import (
	"encoding/json"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/repository/repo"
)

type Refund struct{}

var repoRefund = repo.Database{
	Index: entity.EntityBank,
}

func (f *Refund) FindByReservationId(id string) *string {
	da, err := repoBank.RepoFindByWord("reservedId", id)
	if err != nil {
		return nil
	}

	result := da.Hits[0].Id_

	return &result
}

func (r *Refund) UpdateRefund(id string) {
	doc := json.RawMessage(`{"bankAccount": { "refundState": true }}`)
	repoBank.RepoUpdating(*r.FindByReservationId(id), doc)
}
