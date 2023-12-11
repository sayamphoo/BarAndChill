package wrapper

import (
	"regexp"
	"sayamphoo/microservice/models/entity"
	"unicode"
)

type RegisterWrapper struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RegisterWrapper) CheckRequired() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// ตรวจสอบว่ารายการที่จำเป็นมีค่าหรือไม่
	requiredFields := r.Name != "" &&
		len(r.Phone) == 10 &&
		r.Gender != "" &&
		r.Birthday != "" &&
		emailRegex.MatchString(r.Username) &&
		len(r.Password) >= 6 && len(r.Password) <= 20

	// ตรวจสอบว่ารหัสผ่านประกอบด้วยตัวอักษรหรือไม่
	containsLetters := false
	for _, char := range r.Password {
		if unicode.IsLetter(char) {
			containsLetters = true
			break
		}
	}

	return requiredFields && containsLetters
}

func (data *RegisterWrapper) ToMemberEntity(id string) *entity.MemberEntity {
	return &entity.MemberEntity{
		Id:       id,
		Name:     data.Name,
		Phone:    data.Phone,
		Gender:   data.Gender,
		Birthday: data.Birthday,
		Password: data.Password,
	}
}

type LoginWrapper struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginWrapper) CheckRequired() bool {

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	requiredFields := (emailRegex.MatchString(r.Username) &&
		len(r.Password) >= 6 && len(r.Password) <= 20)
	containsLetters := false
	for _, char := range r.Password {
		if unicode.IsLetter(char) {
			containsLetters = true
			break
		}
	}

	return requiredFields && containsLetters
}

type UpdatePersonalInfoWrapper struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Username string `json:"username"`
}

func (u *UpdatePersonalInfoWrapper) CheckRequired() bool {
	return u.Name != "" &&
		u.Phone != "" &&
		u.Gender != ""
}

type UpdatePassword struct {
	OldPassword string `json:"oldPassword" required:"true"`
	NewPassword string `json:"newPassword" required:"true"`
}

func (u *UpdatePassword) CheckRequired() bool {
	return u.OldPassword != "" &&
		u.NewPassword != ""
}
