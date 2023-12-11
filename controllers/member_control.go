package controllers

import (
	"net/http"
	"sayamphoo/microservice/enum"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/service"

	"github.com/gin-gonic/gin"
)

var serviceMember *service.MemberService

func init() {
	serviceMember = &service.MemberService{}
}

type MemberController struct {
}

func (mcl *MemberController) Register(c *gin.Context) {
	wp := wrapper.RegisterWrapper{}
	c.ShouldBindJSON(&wp)

	if check := wp.CheckRequired(); !check {
		panic(domain.UtilityModel{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	passport := serviceMember.Register(&wp)
	c.JSON(http.StatusOK, passport)
}

func (mcl *MemberController) Login(c *gin.Context) {
	wp := wrapper.LoginWrapper{}
	c.ShouldBindJSON(&wp)

	// if check := wp.CheckRequired(); !check {
	// 	panic(domain.UtilityModel{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "Bad Request",
	// 	})
	// }

	passport := serviceMember.Login(&wp)
	c.JSON(http.StatusOK, passport)
}

func (mcl *MemberController) PersonalInfo(c *gin.Context) {
	id := c.GetString(enum.REQUEST_USER_ID)
	entity := serviceMember.PersonalInfo(id)
	c.JSON(http.StatusOK, entity)
}

func (mcl *MemberController) UpdatePersonalInfo(c *gin.Context) {
	id := c.GetString(enum.REQUEST_USER_ID)

	wp := wrapper.UpdatePersonalInfoWrapper{}
	c.ShouldBindJSON(&wp)

	serviceMember.UpdatePersonalInfo(id, &wp)

	c.JSON(http.StatusOK, domain.UtilityModel{
		Code:    http.StatusOK,
		Message: "Success",
	})

}

func (mcl *MemberController) UpdatePass(c *gin.Context) {
	id := c.GetString(enum.REQUEST_USER_ID)

	wp := wrapper.UpdatePassword{}
	c.ShouldBindJSON(&wp)

	if check := wp.CheckRequired(); !check {
		panic(domain.UtilityModel{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	c.ShouldBindJSON(&wp)

	serviceMember.UpdatePassword(id, &wp)

	c.JSON(http.StatusOK, domain.UtilityModel{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

//------------OWNER------

func (mcl *MemberController) GetAllUser(c *gin.Context) {
	result := serviceMember.GetUser("ALL")
	c.JSON(http.StatusOK, result)
}
