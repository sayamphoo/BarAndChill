package repository

import (
	"encoding/json"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository/repo"
)

type ReservationTableRepo struct{}

var repoReservationTable = repo.Database{
	Index: entity.EntityReserveTableName,
}

func (m *ReservationTableRepo) Sava(wp *wrapper.TableReservationWrapper) (string, error) {
	return repoReservationTable.RepoSave(wp)
}

func (m *ReservationTableRepo) FindTableByDate(date string) (*[]entity.ReserveTable, error) {
	raw, err := repoReservationTable.RepoFindByWord("arrival", date)
	if err != nil {
		return nil, err
	}

	entity, err := entity.HitsToReserveTable(*raw)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m *ReservationTableRepo) findTableByField(field string, value string) (*[]entity.ReserveTable, error) {
	raw, err := repoReservationTable.RepoFindByWord(field, value)
	if err != nil {
		return nil, err
	}

	entity, err := entity.HitsToReserveTable(*raw)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m *ReservationTableRepo) FindTableByIdUser(id string) (*[]entity.ReserveTable, error) {
	return m.findTableByField("user_id", id)
}

func (m *ReservationTableRepo) FindByIdTable(idTable string) (*[]entity.ReserveTable, error) {
	return m.findTableByField("table_id", idTable)
}

func (m *ReservationTableRepo) FindTableById(id string) (*[]entity.ReserveTable, error) {
	return m.findTableByField("_id", id)
}

func (m *ReservationTableRepo) GetAll() (*[]entity.ReserveTable, error) {
	raw, err := repoReservationTable.RepoGetIndex()
	if err != nil {
		return nil, err
	}

	entity, _ := entity.HitsToReserveTable(*raw)
	return entity, nil
}

func (m *ReservationTableRepo) UpdateReserve(reserveId string, doc *json.RawMessage) error {
	_, err := repoReservationTable.RepoUpdating(reserveId, doc)
	return err
}
