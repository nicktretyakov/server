package notes

import (
	"github.com/google/uuid"

	"be/internal/model"
)

type EmailWithUUID struct {
	Email string
	UUID  uuid.UUID
}

func getEmailsWithIDs(users []*model.User, ctxUser *model.User, withEmail bool) (emailsWithIDs []*EmailWithUUID) {
	emailsWithIDs = make([]*EmailWithUUID, 0, len(users))

	for _, user := range mergeRow(users) {
		if !currentUserIsMe(user, ctxUser, withEmail) {
			emailsWithIDs = append(emailsWithIDs, &EmailWithUUID{
				Email: user.Email,
				UUID:  user.ID,
			})
		}
	}

	return
}

func currentUserIsMe(currentUser, ctxUser *model.User, withEmail bool) bool {
	if withEmail {
		if userIsExist(currentUser) && currentUser.Email != "" {
			return ctxUser.ID == currentUser.ID
		}
	}

	if userIsExist(currentUser) {
		return ctxUser.ID == currentUser.ID
	}

	return true
}

func userIsExist(currentUser *model.User) bool {
	return currentUser != nil && currentUser.ID != uuid.Nil
}

func mergeRow(users []*model.User) []*model.User {
	mergeUsers := make([]*model.User, 0, len(users))

	for i, user := range users {
		if i == 0 {
			mergeUsers = append(mergeUsers, user)

			continue
		}

		if !compareUserID(mergeUsers, user) {
			mergeUsers = append(mergeUsers, user)
		}
	}

	return mergeUsers
}

func compareUserID(users []*model.User, baseUser *model.User) bool {
	for _, user := range users {
		if user.ID == baseUser.ID {
			return true
		}
	}

	return false
}

//nolint:exhaustive,gocyclo,cyclop
func GetWhomBookingRecipients(
	notifyType model.NoteEvent,
	booking *model.Booking,
	admin []*model.User,
	ctxUser *model.User,
	email bool,
) []*EmailWithUUID {
	switch notifyType {
	case model.OnRegisterNotify:
		return getEmailsWithIDs(append([]*model.User{booking.Supervisor, booking.Author}, admin...), ctxUser, email)
	case model.OnAgreeNotify:
		return getEmailsWithIDs(append([]*model.User{booking.Supervisor, booking.Assignee}, admin...), ctxUser, email)
	case model.ConfirmedNotify:
		return getEmailsWithIDs(
			append([]*model.User{booking.Supervisor, booking.Author, booking.Assignee}, getParticipants(booking.Participants)...),
			ctxUser,
			email,
		)
	case model.DeclinedNotify:
		return getEmailsWithIDs([]*model.User{booking.Supervisor, booking.Author}, ctxUser, email)
	case model.DoneNotify:
		users := make([]*model.User, 0, len(admin))

		users = append(users, booking.Supervisor, booking.Author, booking.Assignee)
		users = append(users, getParticipants(booking.Participants)...)
		users = append(users, admin...)

		return getEmailsWithIDs(users, ctxUser, email)
	case model.NotSendReportNotify:
		return getEmailsWithIDs([]*model.User{booking.Supervisor}, ctxUser, email)
	case model.MissedReportNotify:
		return getEmailsWithIDs(admin, ctxUser, email)
	case model.SentReportNotify:
		return getEmailsWithIDs(append([]*model.User{booking.Supervisor, booking.Assignee}, admin...), ctxUser, email)
	case model.FinalReportOnRegisterNotify:
		return getEmailsWithIDs(append([]*model.User{booking.Supervisor}, admin...), ctxUser, email)
	case model.FinalReportOnAgreeNotify:
		return getEmailsWithIDs(append([]*model.User{booking.Supervisor, booking.Assignee}, admin...), ctxUser, email)
	case model.FinalReportDeclinedNotify:
		return getEmailsWithIDs([]*model.User{booking.Supervisor, booking.Author}, ctxUser, email)
	default:
		return []*EmailWithUUID{}
	}
}

//nolint:exhaustive
func GetWhomRoomRecipientWithEmail(
	notifyType model.NoteEvent,
	room *model.Room,
	admin []*model.User,
	ctxUser *model.User,
	email bool,
) []*EmailWithUUID {
	switch notifyType {
	case model.OnRegisterNotify:
		return getEmailsWithIDs(append([]*model.User{room.Owner, room.Author}, admin...), ctxUser, email)
	case model.OnAgreeNotify:
		return getEmailsWithIDs(append([]*model.User{room.Owner, room.Employee}, admin...), ctxUser, email)
	case model.ConfirmedNotify:
		users := make([]*model.User, 0, len(admin))

		users = append(users, room.Owner, room.Author, room.Employee)
		users = append(users, getParticipants(room.Participants)...)
		users = append(users, admin...)

		return getEmailsWithIDs(users, ctxUser, email)
	case model.DeclinedNotify:
		return getEmailsWithIDs([]*model.User{room.Owner, room.Author}, ctxUser, email)
	default:
		return []*EmailWithUUID{}
	}
}

func getParticipants(participants []model.Participant) []*model.User {
	users := make([]*model.User, 0, len(participants))

	for _, participant := range participants {
		user := participant.User

		users = append(users, &user)
	}

	return users
}
