package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type User struct {
	ID                 *uuid.UUID  `yaml:"id" db:"id" json:"id"`
	CreatedAt          *time.Time  `yaml:"created_at" db:"created_at" json:"created_at"`
	UpdatedAt          *time.Time  `yaml:"updated_at" db:"updated_at" json:"updated_at"`
	PortalCode         *uint64     `yaml:"portal_code" db:"portal_code" json:"portal_code"`
	ProfileID          *string     `yaml:"profile_id" db:"profile_id" json:"profile_id"`
	Email              *string     `yaml:"email" db:"email" json:"email"`
	Phone              *string     `yaml:"phone" db:"phone" json:"phone"`
	Role               *model.Role `yaml:"role" db:"role" json:"role"`
	EmployeeFirstName  *string     `yaml:"employee_first_name" db:"employee_first_name" json:"employee_first_name"`
	EmployeeMiddleName *string     `yaml:"employee_middle_name" db:"employee_middle_name" json:"employee_middle_name"`
	EmployeePosition   *string     `yaml:"employee_position" db:"employee_position" json:"employee_position"`
	EmployeeLastName   *string     `yaml:"employee_last_name" db:"employee_last_name" json:"employee_last_name"`
	EmployeeAvatar     *string     `yaml:"employee_avatar" db:"employee_avatar" json:"employee_avatar"`
	EmployeeEmail      *string     `yaml:"employee_email" db:"employee_email" json:"employee_email"`
	EmployeePhone      *string     `yaml:"employee_phone" db:"employee_phone" json:"employee_phone"`
}

func (u User) ToModel() model.User {
	return model.User{
		ID:        *u.ID,
		CreatedAt: *u.CreatedAt,
		UpdatedAt: *u.UpdatedAt,
		ProfileID: *u.ProfileID,
		Email:     *u.Email,
		Phone:     *u.Phone,
		Role:      *u.Role,
		Employee: model.Employee{
			FirstName:  u.EmployeeFirstName,
			MiddleName: u.EmployeeMiddleName,
			Position:   u.EmployeePosition,
			LastName:   u.EmployeeLastName,
			Avatar:     u.EmployeeAvatar,
			Email:      u.EmployeeEmail,
			Phone:      u.EmployeePhone,
			PortalCode: *u.PortalCode,
		},
	}
}

func (u User) ToModelPtr() *model.User {
	cu := u.ToModel()
	return &cu
}

func UserFromModel(u model.User) User {
	return User{
		ID:                 &u.ID,
		CreatedAt:          &u.CreatedAt,
		UpdatedAt:          &u.UpdatedAt,
		PortalCode:         &u.Employee.PortalCode,
		ProfileID:          &u.ProfileID,
		Email:              &u.Email,
		Phone:              &u.Phone,
		Role:               &u.Role,
		EmployeeFirstName:  u.Employee.FirstName,
		EmployeeMiddleName: u.Employee.MiddleName,
		EmployeePosition:   u.Employee.Position,
		EmployeeLastName:   u.Employee.LastName,
		EmployeeAvatar:     u.Employee.Avatar,
		EmployeeEmail:      u.Employee.Email,
		EmployeePhone:      u.Employee.Phone,
	}
}

func UserFromModelPtr(u model.User) *User {
	user := UserFromModel(u)
	return &user
}

type UserList []User

func (l UserList) Users() []model.User {
	modelsList := make([]model.User, 0, len(l))
	for _, user := range l {
		modelsList = append(modelsList, user.ToModel())
	}

	return modelsList
}

func (l UserList) UsersPtr() []*model.User {
	modelsList := make([]*model.User, 0, len(l))

	for _, user := range l {
		ptr := user.ToModel()
		modelsList = append(modelsList, &ptr)
	}

	return modelsList
}

func UserListFromModel(users []model.User) UserList {
	res := make(UserList, 0, len(users))

	for _, user := range users {
		res = append(res, UserFromModel(user))
	}

	return res
}
