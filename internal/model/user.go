package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	Unknown Role = iota
	Regular      // Обычный пользователь
	Admin        // РПО
	CEO          // Генеральный директор
	Partner      // Управляющий партнёр
)

type User struct {
	ID        uuid.UUID `yaml:"id"`
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
	ProfileID string    `yaml:"profile_id"` // Id в Цифровом профиле
	Email     string    `yaml:"email"`      // Email пользователя для логина в Цифровом профиле
	Phone     string    `yaml:"phone"`      // Phone пользователя для логина в Цифровом профиле
	Role      Role      `yaml:"role"`       // Роль пользователя
	Employee  Employee
}

func (u User) GetFullName() string {
	return u.Employee.GetLastName() + " " + u.Employee.GetFirstName() + " " + u.Employee.GetMiddleName()
}

type Employee struct {
	FirstName  *string `yaml:"first_name"`
	MiddleName *string `yaml:"middle_name"`
	LastName   *string `yaml:"last_name"`
	Avatar     *string `yaml:"avatar"`
	Email      *string `yaml:"email"`
	Phone      *string `yaml:"phone"`
	Position   *string `yaml:"position"`
	PortalCode uint64  `yaml:"portal_code"`
}

func (e Employee) GetFirstName() string {
	if e.FirstName == nil {
		return ""
	}

	return *e.FirstName
}

func (e Employee) GetMiddleName() string {
	if e.MiddleName == nil {
		return ""
	}

	return *e.MiddleName
}

func (e Employee) GetLastName() string {
	if e.LastName == nil {
		return ""
	}

	return *e.LastName
}

func (e Employee) GetAvatar() string {
	if e.Avatar == nil {
		return ""
	}

	return *e.Avatar
}

func (e Employee) GetEmail() string {
	if e.Email == nil {
		return ""
	}

	return *e.Email
}

func (e Employee) GetPhone() string {
	if e.Phone == nil {
		return ""
	}

	return *e.Phone
}

func (e Employee) GetPosition() string {
	if e.Position == nil {
		return ""
	}

	return *e.Position
}

func (e Employee) GetAvatarLink() string {
	return strings.Join([]string{"/api/v1/avatar", strconv.FormatUint(e.PortalCode, 10)}, "/")
}
