package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"be/internal/model"
)

type (
	Release struct {
		ID          uuid.UUID      `db:"id" yaml:"id"`
		RoomID   uuid.UUID      `db:"room_id" yaml:"room_id"`
		Title       string         `db:"title" yaml:"title"`
		Description string         `db:"description" yaml:"description"`
		CreatedAt   time.Time      `db:"created_at" yaml:"created_at"`
		Date        time.Time      `db:"date" yaml:"date"`
		FactSlot  pgtype.Numeric `db:"fact_slot" yaml:"fact_slot"`
	}

	ReleaseList []Release
)

func ReleaseFromModel(r model.Release) Release {
	factSlot := pgtype.Numeric{}
	_ = factSlot.Set(r.FactSlot.String())

	return Release{
		ID:          r.ID,
		RoomID:   r.RoomID,
		Title:       r.Title,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		Date:        r.Date,
		FactSlot:  factSlot,
	}
}

func (r Release) ToModel() model.Release {
	return model.Release{
		ID:          r.ID,
		RoomID:   r.RoomID,
		Title:       r.Title,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		Date:        r.Date,
		FactSlot:  ToNotification(r.FactSlot),
	}
}

func ToReleaseList(releases []model.Release) ReleaseList {
	releaseList := make(ReleaseList, 0, len(releases))
	for _, release := range releases {
		releaseList = append(releaseList, ReleaseFromModel(release))
	}

	return releaseList
}

func (r ReleaseList) Releases() []model.Release {
	releases := make([]model.Release, 0, len(r))
	for _, release := range r {
		releases = append(releases, release.ToModel())
	}

	return releases
}
