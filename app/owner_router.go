package app

import (
	"sayamphoo/microservice/service"

	"github.com/gin-gonic/gin"
)

func ownerRouter(router *gin.Engine) {

	ownerService := service.OwnerService{}

	v2 := router.Group("api-owner")
	{
		v2.POST("login", ownerService.Login)
		v2.PUT("confirm", tableController.ConfirmReserve)
		v2.GET("get-all-user", memberController.GetAllUser)
		v2.GET("detail-reservation/:id", tableController.GetDetailReservation)
		v2.GET("get-customer-cancel", tableController.GetCustomerCancel)
		v2.PUT("refund-confirm", ownerService.RefundConfirm)
		v2.PUT("song-play", ownerService.SongPlay)
	}
}
