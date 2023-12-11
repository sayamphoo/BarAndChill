package service

import (
	"fmt"
	"net/http"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository"
	"sayamphoo/microservice/security/crypto"
	"sayamphoo/microservice/utility"
	"strings"
)

var repoMember *repository.MemberRepo

func init() {
	repoMember = repository.NewMemberRepo()
}

type MemberService struct{}

func (member *MemberService) Register(model *wrapper.RegisterWrapper) *domain.Passport {

	model.Password = crypto.Hash(model.Password)
	if _, err := repoMember.FindByUsername(model.Username); err == nil {
		panic(domain.UtilityModel{
			Code:    http.StatusConflict,
			Message: "Username already exists",
		})
	}

	result, err := repoMember.Save(model)
	if err != nil {
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "Server Error",
		})
	}

	passport, _ := utility.BuildPassport(&result.Id)

	return passport
}

// Login handles member login.
func (member *MemberService) Login(model *wrapper.LoginWrapper) *domain.Passport {

	entity, err := repoMember.FindByUsername(model.Username)
	if err != nil {

		panic(domain.UtilityModel{
			Code:    http.StatusNotFound,
			Message: "Username Not Found",
		})
	}

	if !crypto.CheckHash(entity.Password, model.Password) {

		panic(domain.UtilityModel{
			Code:    http.StatusUnauthorized,
			Message: "Password incorrect",
		})
	}

	passport, _ := utility.BuildPassport(&entity.Id)

	return passport
}

func (member *MemberService) PersonalInfo(id string) *entity.MemberEntity {
	entity, err := repoMember.FindById(id)
	if err != nil {

		panic(domain.UtilityModel{
			Code:    http.StatusNotFound,
			Message: "User Not Found",
		})
	}

	return entity
}

func (m *MemberService) UpdatePersonalInfo(id string, wp *wrapper.UpdatePersonalInfoWrapper) {

	entity, err := repoMember.FindById(id)
	if err != nil {
		// User not found
		panic(domain.UtilityModel{
			Code:    http.StatusNotFound,
			Message: "User Not Found",
		})
	}

	// Update fields in wp with data from entity only if they are empty
	if wp.Name == "" {
		wp.Name = entity.Name
	}

	if wp.Phone == "" {
		wp.Phone = entity.Phone
	}

	if wp.Gender == "" {
		wp.Gender = entity.Gender
	}

	// Update the information in the system
	if err := repoMember.Update(id, wp); err != nil {
		// Update error
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "Update error",
		})
	}
}

// UpdatePassword updates a member's password.
func (m *MemberService) UpdatePassword(id string, wp *wrapper.UpdatePassword) {
	entity, err := repoMember.FindById(id)

	if err != nil {
		// User not found
		panic(domain.UtilityModel{
			Code:    http.StatusNotFound,
			Message: "User Not Found",
		})
	}

	if check := crypto.CheckHash(entity.Password, wp.OldPassword); !check {
		// Unauthorized
		panic(domain.UtilityModel{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	wp.NewPassword = crypto.Hash(wp.NewPassword)

	update := map[string]string{
		"password": wp.NewPassword,
	}

	if err := repoMember.Update(id, &update); err != nil {
		// Update error
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "Update error",
		})
	}
}

func (m *MemberService) GetUser(id string) *[]entity.MemberEntity {

	if strings.ToUpper(id) == "ALL" {
		result, err := repoMember.GetAllUsers()

		if err != nil {
			return nil
		}

		return result
	} else {
		fmt.Println(id)
		result, err := repoMember.FindById(id)

		if err != nil {
			return nil
		}
		list := []entity.MemberEntity{}
		list = append(list, *result)

		return &list
	}

}
