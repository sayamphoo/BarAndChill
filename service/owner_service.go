package service

import (
	"net/http"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/repository"
	"sayamphoo/microservice/utility"

	"github.com/gin-gonic/gin"
)

type OwnerService struct{}

var repoRefund *repository.Refund

func init() {
	repoRefund = &repository.Refund{}
}

func (O *OwnerService) Login(c *gin.Context) {

	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var logins login
	c.ShouldBindJSON(&logins)

	if logins.Username == "owner" && logins.Password == "123456789" {
		owner := "owner"
		passport, err := utility.BuildPassport(&owner)
		if err != nil {
			panic(domain.UtilityModel{
				Code:    http.StatusNotFound,
				Message: "You not a admin.",
			})
		}

		c.JSON(200, passport)
		return
	}

	c.JSON(http.StatusOK, domain.UtilityModel{
		Code:    404,
		Message: "Username or Password Is Incorrect",
	})

}

func (o *OwnerService) RefundConfirm(c *gin.Context) {
	type RefundConfirm struct {
		ID string `json:"reservationID"`
	}

	var refundConfirm RefundConfirm

	c.ShouldBindJSON(&refundConfirm)
	if chack := repoRefund.FindByReservationId(refundConfirm.ID); chack == nil {
		panic(domain.UtilityModel{
			Code:    http.StatusNotFound,
			Message: "This booking code does not exist.",
		})
	}
	repoRefund.UpdateRefund(refundConfirm.ID)

	c.JSON(http.StatusOK, domain.UtilityModel{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

func (o *OwnerService) SongPlay(c *gin.Context) {
	type RefundConfirm struct {
		ID string `json:"songID"`
	}

	var refundConfirm RefundConfirm
	c.ShouldBindJSON(&refundConfirm)
	repoSong.SetSongState(refundConfirm.ID)

	c.JSON(http.StatusOK, domain.UtilityModel{
		Code:    http.StatusOK,
		Message: "Success",
	})
}
