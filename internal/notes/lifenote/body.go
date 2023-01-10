package lifenote

import (
	"fmt"
	"strings"

	"be/internal/model"
	"be/internal/notes"
)

func getOnRegisterNotify(addressTitle, addressType, shortLink string) *Notice {
	return &Notice{
		Body:        fmt.Sprintf("%s %s добавлен и ожидает регистрации руководителем проектного офиса.", addressType, model.LinkPattern),
		ForEntities: []model.ForEntity{{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink}},
	}
}

func getOnAgreeNotify(addressTitle, addressType, shortLink string, coordinator *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("%s %s зарегистрирован руководителем проектного офиса и ожидает согласования. Согласующий: %s %s",
			addressType,
			model.LinkPattern,
			notes.GetFirstName(coordinator),
			notes.GetLastName(coordinator),
		),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getConfirmedNotify(addressTitle, addressType, shortLink string, actor *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("%s %s утвержден пользователем %s %s.",
			addressType,
			model.LinkPattern,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getDeclinedNotify(addressTitle, addressType, shortLink string, actor *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("%s %s отклонен пользователем %s %s.",
			addressType,
			model.LinkPattern,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getDoneNotify(addressTitle, link string) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Финальный отчет проекта согласован. Проект %s завершен. Спасибо за работу!", model.LinkPattern),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: link},
		},
	}
}

func getNotSendReportNotify(addressTitle, shortLink string, periods []string) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Добавьте статусный отчет по проекту %s за следующие месяцы: %s.", model.LinkPattern, strings.Join(periods, ", ")),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getMissedReportNotify(data []*notes.MissedReportNotifyData) *Notice {
	var (
		msg         = "Не добавлены статусные отчеты по проектам:\n"
		forEntities = make([]model.ForEntity, 0, len(data))
	)

	for _, d := range data {
		newPart := fmt.Sprintf("%s (%s)\n", d.BookingTitle, model.LinkPattern)

		forEntities = append(forEntities, model.ForEntity{Pattern: model.LinkPattern, Value: d.BookingTitle, Link: d.BookingLink})

		for _, period := range d.Periods {
			newPart = fmt.Sprintf("%s - отчет за %s\n", newPart, period)
		}

		msg += fmt.Sprintf("%s Руководитель проекта: %s %s\n", newPart, notes.GetFirstName(d.Rpo), notes.GetLastName(d.Rpo))
	}

	return &Notice{
		Body:        msg,
		ForEntities: forEntities,
	}
}

func getSentReportNotify(addressTitle, shortLink string, actor *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Отчет по проекту %s добавлен пользователем %s %s.",
			model.LinkPattern,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getFinalReportOnRegisterNotify(addressTitle, shortLink string) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Финальный отчет проекта %s ожидает регистрации руководителем проектного офиса.", model.LinkPattern),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getFinalReportOnAgreeNotify(addressTitle, shortLink string, coordinator *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Финальный отчет проекта %s зарегистрирован Руководителем проектного офиса и ожидает согласования. Согласующий: %s %s.",
			model.LinkPattern,
			notes.GetFirstName(coordinator),
			notes.GetLastName(coordinator),
		),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}

func getFinalReportDeclinedNotify(addressTitle, shortLink string, actor *model.User) *Notice {
	return &Notice{
		Body: fmt.Sprintf("Финальный отчет проекта %s отклонен пользователем %s %s.",
			model.LinkPattern,
			notes.GetFirstName(actor),
			notes.GetLastName(actor)),
		ForEntities: []model.ForEntity{
			{Pattern: model.LinkPattern, Value: addressTitle, Link: shortLink},
		},
	}
}
