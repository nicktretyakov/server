package model

import (
	"io"
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID             uuid.UUID
	AuthorID       uuid.UUID
	CreatedAt      time.Time
	AddressID uuid.UUID

	Key      string
	Filename string
	Mime     string
	Size     int

	source io.Reader
	url    string
}

func (a *Attachment) Source() io.Reader {
	return a.source
}

func (a *Attachment) SetSource(r io.Reader) {
	a.source = r
}

func (a *Attachment) URL() string {
	return a.url
}

func (a *Attachment) SetURL(url string) {
	a.url = url
}
