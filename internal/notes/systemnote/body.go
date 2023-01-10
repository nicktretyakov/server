package systemnote

import (
	"fmt"
	"strings"

	"be/internal/model"
	"be/internal/notes"
)

func getOnRegisterNotify(addressTitle, addressType string) *Notice {
	return &Notice{
		Header: addressTitle,
		Body:   fmt.Sprintf("%s %s добавлен и ожидает регистрации руководителем проектного офиса.", addressType, addressTitle),
	}
}

func getOnAgreeNotify(addressTitle, addressType string, coordinator *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("%s %s зарегистрирован руководителем проектного офиса и ожидает согласования. Согласующий: %s %s",
			addressType,
			addressTitle,
			notes.GetFirstName(coordinator),
			notes.GetLastName(coordinator),
		),
	}
}

func getConfirmedNotify(addressTitle, addressType string, actor *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("%s %s утвержден пользователем %s %s.",
			addressType,
			addressTitle,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
	}
}

func getDeclinedNotify(addressTitle, addressType string, actor *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("%s %s отклонен пользователем %s %s.",
			addressType,
			addressTitle,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
	}
}

func getDoneNotify(addressTitle string) *Notice {
	return &Notice{
		Header: addressTitle,
		Body:   fmt.Sprintf("Финальный отчет проекта согласован. Проект %s завершен. Спасибо за работу!", addressTitle),
	}
}

func getNotSendReportNotify(addressTitle string, periods []string) *Notice {
	return &Notice{
		Header: addressTitle,
		Body:   fmt.Sprintf("Добавьте статусный отчет по проекту %s за следующие месяцы: %s.", addressTitle, strings.Join(periods, ", ")),
	}
}

func getMissedReportNotify(data []*notes.MissedReportNotifyData) *Notice {
	msg := "<p> Не добавлены статусные отчеты по проектам:<br><br>"

	for _, d := range data {
		newPart := fmt.Sprintf("<br> %s <br>", d.BookingLink)

		for _, period := range d.Periods {
			newPart = fmt.Sprintf("%s - отчет за %s<br>", newPart, period)
		}

		msg += fmt.Sprintf("%s Руководитель проекта: %s %s<br>", newPart, notes.GetFirstName(d.Rpo), notes.GetLastName(d.Rpo))
	}

	msg += " </p>"

	return &Notice{
		Header: "Проекты, ожидающие отчетов",
		Body:   msg,
	}
}

func getSentReportNotify(addressTitle string, actor *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("Отчет по проекту %s добавлен пользователем %s %s.",
			addressTitle,
			notes.GetFirstName(actor),
			notes.GetLastName(actor),
		),
	}
}

func getFinalReportOnRegisterNotify(addressTitle string) *Notice {
	return &Notice{
		Header: addressTitle,
		Body:   fmt.Sprintf("Финальный отчет проекта %s ожидает регистрации руководителем проектного офиса.", addressTitle),
	}
}

func getFinalReportOnAgreeNotify(addressTitle string, coordinator *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("Финальный отчет проекта %s зарегистрирован Руководителем проектного офиса и ожидает согласования. Согласующий: %s %s.",
			addressTitle,
			notes.GetFirstName(coordinator),
			notes.GetLastName(coordinator),
		),
	}
}

func getFinalReportDeclinedNotify(addressTitle string, actor *model.User) *Notice {
	return &Notice{
		Header: addressTitle,
		Body: fmt.Sprintf("Финальный отчет проекта %s отклонен пользователем %s %s.",
			addressTitle,
			notes.GetFirstName(actor),
			notes.GetLastName(actor)),
	}
}
