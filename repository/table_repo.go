package repository

import (
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository/repo"
)

type TableRepo struct{}

var repoTable = repo.Database{
	Index: entity.EntityTableName,
}

func (m *TableRepo) Sava(wp *wrapper.TableWrapper) (string, error) {
	id, err := repoTable.RepoSave(wp)
	if err != nil {
		return "", err
	}
	return id, nil
}
