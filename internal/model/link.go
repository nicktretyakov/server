package model

import "github.com/google/uuid"

//nolint:stylecheck
type Link struct {
	Id     uuid.UUID
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
}
