package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sayamphoo/microservice/enum"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/service"
	"sayamphoo/microservice/utility"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var serviceTable *service.TableService

func init() {
	serviceTable = &service.TableService{}
}

type TableController struct{}

func (t *TableController) GetTable(c *gin.Context) {
	date := c.Param("date")
	table := serviceTable.GetTable(date)

	c.JSON(200, table)
}

func (t *TableController) Reservation(c *gin.Context) {

	file, err := c.FormFile("statement")
	if err != nil {
		panic(domain.UtilityModel{
			Code:    http.StatusBadRequest,
			Message: "Missing or invalid 'statement' file in the request",
		})
	}

	fileName := fmt.Sprintf("%s.jpg", time.Now().Format("20060102150405"))
	dir := filepath.Join(os.Getenv("RESOURCE_PATH"), fileName)
	c.SaveUploadedFile(file, dir)

	userID := c.GetString(enum.REQUEST_USER_ID)
	drinkID := c.PostForm("drink_id")
	tableID := c.PostForm("table_id")
	arrival := c.PostForm("arrival")

	if userID == "" || drinkID == "" || tableID == "" || arrival == "" {
		panic(domain.UtilityModel{
			Code:    http.StatusBadRequest,
			Message: "One or more required fields are missing",
		})
	}

	member := serviceMember.GetUser(userID)
	if member == nil {
		c.JSON(404, "Member Not Found")
		return
	}

	table := serviceTable.ReservationChackIdTable(tableID, arrival)
	if table {
		c.JSON(409, "Already reserved")
		return
	}

	wp := wrapper.TableReservationWrapper{
		UserID:    userID,
		DrinkID:   drinkID,
		TableID:   tableID,
		Arrival:   arrival,
		Status:    "waiting",
		Timestamp: utility.GetTimeNow(),
		Statement: fileName,
	}

	_, err = serviceTable.Reservation(&wp)
	if err != nil {
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "เกิดข้อผิดพลาด",
		})
	}

	c.JSON(200, "Confirm")
}

func (t *TableController) GetMyReservation(c *gin.Context) {
	id := c.GetString(enum.REQUEST_USER_ID)
	date := c.Param("date")
	entity, _ := serviceTable.GetMyReservation(id, date)
	c.JSON(200, entity)
}

func (t *TableController) Refund(c *gin.Context) {

	id := c.GetString(enum.REQUEST_USER_ID)
	var wrapper wrapper.RefundWrapper

	if err := c.ShouldBindJSON(&wrapper); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	wrapper.Timestamp = utility.GetTimeNow()
	if err := serviceTable.Refund(id, &wrapper); err != nil {
		c.JSON(http.StatusNotFound, "User or Reservation Not Found")
	}

	c.JSON(http.StatusOK, "")

}

//----------onwer

func (t *TableController) GetDetailReservation(c *gin.Context) {
	id := c.Param("id")
	data, err := serviceTable.GetDetailReservation(id)

	if err != nil {
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "Failed",
		})
	}
	result := make([]domain.ReserveTableDomain, 0)
	for _, x := range *data {
		name := serviceMember.GetUser(strings.TrimSpace(x.UserID))
		userName := ""
		if name != nil {
			userName = (*name)[0].Name
		}
		result = append(result, domain.ReserveTableDomain{
			ID:        x.ID,
			UserID:    x.UserID,
			NameUser:  userName,
			DrinkID:   x.DrinkID,
			TableID:   x.TableID,
			Arrival:   x.Arrival,
			Timestamp: x.Timestamp,
			Statement: x.Statement,
			Status:    x.Status,
		})
	}

	c.JSON(200, result)
}

func (t *TableController) ConfirmReserve(c *gin.Context) {
	type ReserveRequest struct {
		ReserveID string `json:"reserve_id"`
		State     bool   `json:"state"`
	}

	var request ReserveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	data := serviceTable.ConfirmReserve(request.ReserveID, request.State)
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Reservation confirmed"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to confirm reservation"})
	}
}

func (m *TableController) GetCustomerCancel(c *gin.Context) {

	domain := serviceTable.GetCustomerCancel()
	c.JSON(http.StatusOK, domain)

}
