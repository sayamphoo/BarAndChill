package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository"
	"strings"
)

var repoReservationTable *repository.ReservationTableRepo
var repoBank *repository.BankRepo

func init() {
	repoReservationTable = &repository.ReservationTableRepo{}
	repoBank = &repository.BankRepo{}
}

type TableService struct {
}

func (t *TableService) GetTable(date string) *[]domain.TableDomain {

	table := entity.GetTable()
	reservation, _ := repoReservationTable.FindTableByDate(date)

	domainTable := make([]domain.TableDomain, 0)

	for i := 0; i < len(table); i++ {
		d := domain.TableDomain{
			Id:    table[i].ID,
			Name:  table[i].Name,
			Chair: table[i].Chair,
			State: false,
		}

		if reservation != nil {
			for j := 0; j < len(*reservation); j++ {
				if table[i].ID == (*reservation)[j].TableID {
					if strings.ToLower(((*reservation)[j].Status)[:6]) == "cancel" { //cancel
						d.State = false
					} else {
						d.State = true
					}
					break
				}
			}
		}

		domainTable = append(domainTable, d)
	}

	return &domainTable
}

func (t *TableService) Reservation(wp *wrapper.TableReservationWrapper) (*domain.UtilityModel, error) {
	_, err := repoReservationTable.Sava(wp)
	if err != nil {
		return nil, err
	}

	return &domain.UtilityModel{
		Code:    200,
		Message: "SUCCESSFULLY",
	}, nil
}

func (t *TableService) ReservationChackIdTable(tableId string, arrival string) bool {
	raw, _ := repoReservationTable.FindByIdTable(tableId)

	if raw != nil {
		for _, x := range *raw {
			if arrival == x.Arrival {
				return true
			}
		}
	}

	return false
}
func (t *TableService) GetMyReservation(id string, date string) ([]entity.ReserveTable, error) {
	raw, err := repoReservationTable.FindTableByIdUser(id)
	if err != nil {
		return nil, err
	}

	domain := make([]entity.ReserveTable, 0)

	if strings.ToUpper(date) == "ALL" {
		return *raw, nil
	}

	for _, reservation := range *raw {
		if date == reservation.Arrival {
			d := entity.ReserveTable{
				ID:        reservation.ID,
				TableID:   reservation.TableID,
				Arrival:   reservation.Arrival,
				Status:    reservation.Status,
				UserID:    reservation.UserID,
				DrinkID:   reservation.DrinkID,
				Timestamp: reservation.Timestamp,
				Statement: reservation.Statement,
			}

			domain = append(domain, d)
		}
	}

	return domain, nil
}

func (t *TableService) Refund(idUser string, w *wrapper.RefundWrapper) error {
	raw, err := repoReservationTable.FindTableById(w.ReservedID)

	if err != nil {
		return err
	}

	if (*raw)[0].UserID != idUser && (*raw)[0].Status != "cancel" {
		return errors.New("user and reservation don't match or reservation cancel already")
	}

	doc := json.RawMessage(`{ "status": "cancel-customer" }`)
	if err := repoReservationTable.UpdateReserve(w.ReservedID, &doc); err != nil {
		return err
	}
	w.BankAccount.RefundState = false
	repoBank.Sava(w)

	return nil
}

// ------------ onwer

func (m *TableService) GetDetailReservation(id string) (*[]entity.ReserveTable, error) {
	var raw *[]entity.ReserveTable
	var err error
	if strings.ToUpper(id) == "ALL" {
		raw, err = repoReservationTable.GetAll()
	} else {
		raw, err = repoReservationTable.FindTableById(id)
	}
	return raw, err
}

func (m *TableService) GetCustomerCancel() []domain.ReserveRefundDomain {
	data, err := m.GetDetailReservation("all")

	if err != nil {
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "Failed",
		})
	}

	domains := make([]domain.ReserveRefundDomain, 0)

	for _, element := range *data {
		if element.Status == "cancel-customer" {
			bank := repoBank.FindByIdReservation(element.ID)
			nameUser, _ := repoMember.FindById(element.UserID)
			d := domain.ReserveRefundDomain{
				ID:              element.ID,
				UserID:          element.UserID,
				NameUser:        nameUser.Name,
				DrinkID:         element.DrinkID,
				TableID:         element.TableID,
				Arrival:         element.Arrival,
				Timestamp:       element.Timestamp,
				Statement:       element.Statement,
				Status:          element.Status,
				TimestampCancel: bank.Timestamp,
				Refund:          bank.BankAccount,
			}
			domains = append(domains, d)
		}
	}

	return domains
}

func (m *TableService) ConfirmReserve(reserveId string, state bool) error {
	var doc json.RawMessage
	if state {
		doc = json.RawMessage(`{ "status": "confirm" }`)
	} else {
		doc = json.RawMessage(`{ "status": "cancel-owner" }`)
	}
	if err := repoReservationTable.UpdateReserve(reserveId, &doc); err != nil {
		return err
	}
	return nil
}
