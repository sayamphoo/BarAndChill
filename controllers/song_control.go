package controllers

import (
	"sayamphoo/microservice/enum"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/service"

	"github.com/gin-gonic/gin"
)

var serviceSong *service.SongService

func init() {
	serviceSong = &service.SongService{}
}

type SongController struct{}

func (sc *SongController) Request(c *gin.Context) {
	wp := wrapper.SongRequestWrapper{}
	c.ShouldBindJSON(&wp)

	wp.UserID = c.GetString(enum.REQUEST_USER_ID)
	queue := serviceSong.Request(&wp)

	c.JSON(200, domain.SongDomain{
		Queue: queue,
	})

}

func (sc *SongController) Queuelist(c *gin.Context) {
	entity := serviceSong.GetSong()
	c.JSON(200, entity)
}
