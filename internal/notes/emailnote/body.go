package emailnote

import (
	"fmt"
	"strings"

	"be/internal/model"
	"be/internal/notes"
)

func getOnRegisterNotify(addressTitle, addressType, link string) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     addressType,
				SecondPartBody:    "добавлен и ожидает регистрации руководителем проектного офиса.",
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getOnAgreeNotify(addressTitle, addressType, link string, coordinator *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody: addressType,
				SecondPartBody: fmt.Sprintf(
					"зарегистрирован руководителем проектного офиса и ожидает согласования. Согласующий: %s %s.",
					notes.GetFirstName(coordinator), notes.GetLastName(coordinator)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getConfirmedNotify(addressTitle, addressType, link string, actor *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     addressType,
				SecondPartBody:    fmt.Sprintf("утвержден пользователем %s %s.", notes.GetFirstName(actor), notes.GetLastName(actor)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getDeclinedNotify(addressTitle, addressType, link string, actor *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     addressType,
				SecondPartBody:    fmt.Sprintf("отклонен пользователем %s %s.", notes.GetFirstName(actor), notes.GetLastName(actor)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getDoneNotify(addressTitle, link string, _ *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     "Финальный отчет проекта согласован. Проект ",
				SecondPartBody:    "завершен. Спасибо за работу!",
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getNotSendReportNotify(addressTitle, link string, periods []string) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     "Добавьте статусный отчет по проекту: ",
				SecondPartBody:    fmt.Sprintf(" за следующие месяцы: %s.", strings.Join(periods, ", ")),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

type MissedReportNotifyData struct {
	BookingTitle string
	BookingLink  string
	Periods      []string
	Rpo          *model.User
}

func getMissedReportNotify(data []*MissedReportNotifyData) *Data {
	var secondPart string

	dataBody := make([]Body, 0, len(data))
	msgFirstPart := "Не добавлены статусные отчеты по проектам: {{.Br}}{{.Br}}"

	for _, d := range data {
		secondPart = "{{.Br}}"

		for _, period := range d.Periods {
			secondPart += fmt.Sprintf("- отчет за %s{{.Br}}", period)
		}

		secondPart += fmt.Sprintf(
			"Руководитель проекта: %s %s {{.Br}}{{.Br}}",
			notes.GetFirstName(d.Rpo),
			notes.GetLastName(d.Rpo),
		)

		dataBody = append(dataBody, Body{
			FirstPartBody:     msgFirstPart,
			SecondPartBody:    secondPart,
			AddressTitle: d.BookingTitle,
			Link:              d.BookingLink,
			ForReport:         true,
		})

		msgFirstPart = "{{.Br}}"
	}

	return &Data{DataBody: dataBody}
}

func getSentReportNotify(addressTitle, link string, actor *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     "Отчет по проекту ",
				SecondPartBody:    fmt.Sprintf("добавлен пользователем %s %s.", notes.GetFirstName(actor), notes.GetLastName(actor)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getFinalReportOnRegisterNotify(addressTitle, link string) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     "Финальный отчет проекта ",
				SecondPartBody:    "ожидает регистрации руководителем проектного офиса.",
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getFinalReportOnAgreeNotify(addressTitle, link string, coordinator *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody: "Финальный отчет проекта ",
				SecondPartBody: fmt.Sprintf(
					"зарегистрирован Руководителем проектного офиса и ожидает согласования. Согласующий: %s %s.",
					notes.GetFirstName(coordinator), notes.GetLastName(coordinator)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}

func getFinalReportDeclinedNotify(addressTitle, link string, actor *model.User) *Data {
	return &Data{
		DataBody: []Body{
			{
				FirstPartBody:     "Финальный отчет проекта ",
				SecondPartBody:    fmt.Sprintf("отклонен пользователем %s %s.", notes.GetFirstName(actor), notes.GetLastName(actor)),
				AddressTitle: addressTitle,
				Link:              link,
				ForReport:         false,
			},
		},
	}
}
