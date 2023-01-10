package outmember

import (
	"github.com/google/uuid"

	"be/internal/acl"
	"be/internal/model"
)

type IOutmember interface {
	Type() model.OutmemberType
	AddressID() uuid.UUID
	Source() model.Outmember
	ACL(user model.User, io acl.AddressACL) bool
}

type InitialOutmember struct {
	Outmember model.Outmember
}

func (i InitialOutmember) Result() model.OutmemberResult {
	return i.Outmember.Result
}

func (i InitialOutmember) AddressID() uuid.UUID {
	return i.Outmember.AddressID
}

func (i InitialOutmember) Source() model.Outmember {
	return i.Outmember
}

func (i InitialOutmember) ACL(user model.User, io acl.AddressACL) bool {
	return acl.CanCreateInitialOutmember(user, io)
}

func (i InitialOutmember) Type() model.OutmemberType {
	return i.Outmember.Type
}

type AcceptanceOutmember struct {
	Outmember model.Outmember
}

func (i AcceptanceOutmember) Result() model.OutmemberResult {
	return i.Outmember.Result
}

func (i AcceptanceOutmember) AddressID() uuid.UUID {
	return i.Outmember.AddressID
}

func (i AcceptanceOutmember) Source() model.Outmember {
	return i.Outmember
}

func (i AcceptanceOutmember) ACL(user model.User, io acl.AddressACL) bool {
	return acl.CanCreateAcceptanceOutmember(user, io)
}

func (i AcceptanceOutmember) Type() model.OutmemberType {
	return i.Outmember.Type
}

type RegisterFinalOutmember struct {
	Outmember   model.Outmember
	FinalReport model.FinalReport
}

func (f RegisterFinalOutmember) AddressID() uuid.UUID {
	return f.Outmember.AddressID
}

func (f RegisterFinalOutmember) Source() model.Outmember {
	return f.Outmember
}

func (f RegisterFinalOutmember) ACL(user model.User, io acl.AddressACL) bool {
	return acl.CanRegisterFinalOutmember(user, io)
}

func (f RegisterFinalOutmember) Result() model.OutmemberResult {
	return f.Outmember.Result
}

func (f RegisterFinalOutmember) Type() model.OutmemberType {
	return f.Outmember.Type
}

type ApproveFinalOutmember struct {
	Outmember   model.Outmember
	FinalReport model.FinalReport
}

func (f ApproveFinalOutmember) AddressID() uuid.UUID {
	return f.Outmember.AddressID
}

func (f ApproveFinalOutmember) Source() model.Outmember {
	return f.Outmember
}

func (f ApproveFinalOutmember) ACL(user model.User, io acl.AddressACL) bool {
	return acl.CanApproveFinalOutmember(user, io)
}

func (f ApproveFinalOutmember) Result() model.OutmemberResult {
	return f.Outmember.Result
}

func (f ApproveFinalOutmember) Type() model.OutmemberType {
	return f.Outmember.Type
}
