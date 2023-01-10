package model

import "github.com/google/uuid"

type BookingUserRole int

const (
	SupervisorBookingUserRole  BookingUserRole = 1 // Создатель проекта
	AssigneeBookingUserRole    BookingUserRole = 2 // Согласующий проекта
	AuthorBookingUserRole      BookingUserRole = 3 // Автор проекта
	ParticipantBookingUserRole BookingUserRole = 4 // Участник проекта
)

type BookingUser struct {
	BookingID       uuid.UUID       `yaml:"booking_id"`
	UserID          uuid.UUID       `yaml:"user_id"`
	Role            BookingUserRole `yaml:"role"`
	RoleDescription string          `yaml:"role_description"`
}
