package model

import (
	"time"

	"github.com/google/uuid"
)

type Outmember struct {
	ID             uuid.UUID              `yaml:"id"`
	CreatedAt      time.Time              `yaml:"created_at"`
	Type           OutmemberType          `yaml:"type"`
	UserID         uuid.UUID              `yaml:"user_id"`
	User           *User                  `yaml:"user"`
	AddressID uuid.UUID              `yaml:"address_id"`
	Result         OutmemberResult        `yaml:"result"`
	Extra          map[string]interface{} `yaml:"extra"`
	Role           OutmemberRole          `yaml:"role"`
}

type (
	OutmemberType   int
	OutmemberResult int
	OutmemberRole   int
)

func (r OutmemberResult) Bool() bool {
	return r == AcceptOutmemberResult
}

const (
	DeclineOutmemberResult OutmemberResult = 0
	AcceptOutmemberResult  OutmemberResult = 1
)

const (
	UnknownOutmemberType       OutmemberType = 0
	InitialOutmemberType       OutmemberType = 1
	FinalRegisterOutmemberType OutmemberType = 2
	ApprovalOutmemberType      OutmemberType = 3
	FinalApprovalOutmemberType OutmemberType = 4
)

const (
	UnknownOutmemberRole              OutmemberRole = 0
	BookingManagerOutmemberRole OutmemberRole = 1
	AssigneeOutmemberRole             OutmemberRole = 2
)
