package model

import (
	"strings"

	"github.com/google/uuid"
)

type (
	Department struct {
		ID    uuid.UUID `db:"id" yaml:"id"`
		Title string    `db:"title" yaml:"title"`
	}

	DepartmentModelList []Department
)

type DepartmentID struct {
	ID uuid.UUID `json:"id"`
}

func (d DepartmentModelList) GetDepartments() string {
	titles := make([]string, 0, len(d))
	for _, item := range d {
		titles = append(titles, item.Title)
	}

	return strings.Join(titles, ", ")
}
