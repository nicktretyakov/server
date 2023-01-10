package notes

import (
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

const (
	AddressTypeBooking = "Бронирование"
	AddressTypeRoom = "Комната"

	ReplaceTitle = "{{.AddressTitle}}"
	ReplaceType  = "{{.AddressType}}"
	ReportPeriod = "{{.ReportPeriod}}"
)

//nolint:gochecknoglobals
var Subjects = map[model.NoteEvent]string{
	model.OnRegisterNotify:            ReplaceType + " " + ReplaceTitle + " добавлен и ожидает регистрации.",
	model.OnAgreeNotify:               ReplaceType + " " + ReplaceTitle + " зарегистрирован и ожидает согласования.",
	model.ConfirmedNotify:             ReplaceType + " " + ReplaceTitle + " утвержден.",
	model.DeclinedNotify:              ReplaceType + " " + ReplaceTitle + " отклонен.",
	model.DoneNotify:                  "Проект " + ReplaceTitle + " завершен.",
	model.NotSendReportNotify:         "Ожидается отчет(ы) по проекту " + ReplaceTitle + ".",
	model.MissedReportNotify:          "Проекты, ожидающие отчетов.",
	model.SentReportNotify:            "Добавлен отчет по проекту " + ReplaceTitle + " за " + ReportPeriod + ".",
	model.FinalReportOnRegisterNotify: "К проекту " + ReplaceTitle + " добавлен финальный отчет.",
	model.FinalReportOnAgreeNotify:    "Финальный отчет проекта " + ReplaceTitle + " зарегистрирован и ожидает согласования.",
	model.FinalReportDeclinedNotify:   "Финальный отчет проекта " + ReplaceTitle + " отклонен.",
}

//nolint:gocyclo,cyclop,nakedret
func YearMonthString(t time.Time) (month string) {
	switch t.Month() {
	case time.January:
		month = "Январь"
	case time.February:
		month = "Февраль"
	case time.March:
		month = "Март"
	case time.April:
		month = "Апрель"
	case time.May:
		month = "Май"
	case time.June:
		month = "Июнь"
	case time.July:
		month = "Июль"
	case time.August:
		month = "Август"
	case time.September:
		month = "Сентябрь"
	case time.October:
		month = "Октябрь"
	case time.November:
		month = "Ноябрь"
	case time.December:
		month = "Декабрь"
	}

	month += " " + strconv.Itoa(t.Year())

	return
}

func GetFirstName(user *model.User) string {
	if user == nil {
		return ""
	}

	if user.Employee.FirstName != nil {
		return *user.Employee.FirstName
	}

	return ""
}

func GetLastName(user *model.User) string {
	if user == nil {
		return ""
	}

	if user.Employee.LastName != nil {
		return *user.Employee.LastName
	}

	return ""
}

type MissedReportNotifyData struct {
	BookingTitle string
	BookingLink  string
	Periods      []string
	Rpo          *model.User
}

type Link struct {
	Base             string
	TypeAddress string
}

func CreateAddressViewLink(addressID uuid.UUID, link Link) string {
	return path.Join(link.Base, strings.Replace(link.TypeAddress, "[[uuid]]", addressID.String(), 1))
}
