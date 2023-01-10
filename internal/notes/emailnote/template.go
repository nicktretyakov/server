package emailnote

import (
	"bytes"
	"html/template"
	"strings"

	"be/internal/model"
	"be/internal/notes"
)

type Data struct {
	DataBody []Body
}

type Body struct {
	FirstPartBody     string
	SecondPartBody    string
	AddressTitle string
	Link              string
	ForReport         bool
}

func executeBody(subject string, data interface{}) (string, error) {
	var (
		body bytes.Buffer
		t    = template.Must(template.New(subject).Parse(HTMLTemplate))
	)

	if err := t.Execute(&body, data); err != nil {
		return "", err
	}

	return strings.ReplaceAll(body.String(), "{{.Br}}", "<br>"), nil
}

func getSubject(notifyType model.NoteEvent, addressTitle, addressType, reportPeriod string) string {
	subject := notes.Subjects[notifyType]
	subject = strings.ReplaceAll(subject, notes.ReplaceTitle, addressTitle)
	subject = strings.ReplaceAll(subject, notes.ReplaceType, addressType)
	subject = strings.ReplaceAll(subject, notes.ReportPeriod, reportPeriod)

	return subject
}
