package app

import (
	"github.com/gin-gonic/gin"
)

func customRouter(router *gin.Engine) {
	v1 := router.Group("api-customer")
	{
		member := v1.Group("secret/member")
		{
			member.POST("register", memberController.Register) //ByPass
			member.POST("login", memberController.Login)       //ByPass
			member.GET("personal-info", memberController.PersonalInfo)
			member.PUT("update-personal-info", memberController.UpdatePersonalInfo)
			member.PUT("update-password", memberController.UpdatePass)
		}

		table := v1.Group("bar/table")
		{
			table.GET("get-table/:date", tableController.GetTable)
			table.POST("reservation", tableController.Reservation)
			table.GET("get-my-reservation/:date", tableController.GetMyReservation)
			table.POST("refund-money", tableController.Refund)
		}

		song := v1.Group("bar/soung")
		{
			song.POST("request", songController.Request)
			song.GET("queue-list", songController.Queuelist)
		}
	}
}
