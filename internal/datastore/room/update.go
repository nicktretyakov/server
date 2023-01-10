package room

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

//nolint:gocognit
func (s *Storage) Update(ctx context.Context, room model.Room, equipmentIDs, slotIDs, bookingIDs []uuid.UUID) (*model.Room, error) {
	roomToStore := dbmodel.RoomFromModel(room)
	roomToStore.UpdatedAt = s.db.Now()
	roomToStore.Bookings = bookingIDs

	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := s.db.ExecTxBuilder(ctx, tx, roomUpdateQuery(roomToStore)); err != nil {
			return err
		}

		if _, err := s.db.ExecTxBuilder(ctx, tx, roomSlotDeleteQuery(roomToStore)); err != nil {
			return err
		}

		if _, err := s.db.ExecTxBuilder(ctx, tx, roomEquipmentDeleteQuery(roomToStore)); err != nil {
			return err
		}

		if len(equipmentIDs) > 0 {
			if _, err := s.db.ExecTxBuilder(ctx, tx, equipmentUpdateQuery(roomToStore.ID, equipmentIDs)); err != nil {
				return err
			}
		}

		if len(slotIDs) > 0 {
			if _, err := s.db.ExecTxBuilder(ctx, tx, slotUpdateQuery(roomToStore.ID, slotIDs)); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return roomToStore.ToModelPtr(), nil
}

func (s *Storage) UpdateStatus(ctx context.Context, roomID uuid.UUID, status model.Status) error {
	room := dbmodel.Room{Status: status, ID: roomID}
	room.UpdatedAt = s.db.Now()

	if cmd, err := s.db.ExecBuilder(ctx, roomUpdateStatusQuery(room)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *Storage) UpdateState(ctx context.Context, roomID uuid.UUID, state model.State) error {
	room := dbmodel.Room{State: state, ID: roomID}
	room.UpdatedAt = s.db.Now()

	if cmd, err := s.db.ExecBuilder(ctx, roomUpdateStateQuery(room)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *Storage) UpdateBookingsRoom(ctx context.Context, roomID uuid.UUID, bookingIDs []uuid.UUID) error {
	room := dbmodel.Room{Bookings: bookingIDs, ID: roomID}
	room.UpdatedAt = s.db.Now()

	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := s.db.ExecTxBuilder(ctx, tx, roomDeleteBookingsQuery(room)); err != nil {
			return err
		}

		if _, err := s.db.ExecTxBuilder(ctx, tx, roomUpdateBookingsQuery(room)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateLinks(ctx context.Context, roomID uuid.UUID, link []model.Link) error {
	room := dbmodel.Room{Links: link, ID: roomID}
	room.UpdatedAt = s.db.Now()

	if cmd, err := s.db.ExecBuilder(ctx, roomUpdateLinksQuery(room)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

type participantDB struct {
	UserID uuid.UUID `db:"user_id"`
	Role   string    `db:"role"`
}

func (s *Storage) UpdateParticipant(ctx context.Context, roomID uuid.UUID, participants map[uuid.UUID]string) error {
	room := dbmodel.Room{
		ID:        roomID,
		UpdatedAt: s.db.Now(),
	}

	if cmd, err := s.db.ExecBuilder(
		ctx,
		roomUpdateParticipantsQuery(room, getParticipants(participants)),
	); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func getParticipants(participants map[uuid.UUID]string) []participantDB {
	participantsDB := make([]participantDB, 0, len(participants))

	for participantID, role := range participants {
		participantsDB = append(participantsDB, participantDB{
			UserID: participantID,
			Role:   role,
		})
	}

	return participantsDB
}

func roomUpdateQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at":      room.UpdatedAt,
			"number":          room.Number,
			"author_id":       room.Author.ID,
			"title":           room.Title,
			"description":     room.Description,
			"target_audience": room.TargetAudience,
			"status":          room.Status,
			"links":           room.Links,
			"employee_id":     room.Employee.ID,
			"owner_id":        room.Owner.ID,
			"booking_ids":     room.Bookings,
			"creation_date":   room.CreationDate,
			"space":           room.Space,
	        "security_email":  room.SecurityEmail,
	        "visible":         room.Visible,
		}).
		Where("id=?", room.ID)
}

func roomSlotDeleteQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(slotTableName).
		Set("room_id", nil).
		Where("room_id=?", room.ID)
}

func roomEquipmentDeleteQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(equipmentTableName).
		Set("room_id", nil).
		Where("room_id=?", room.ID)
}

func roomUpdateStatusQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at": room.UpdatedAt,
			"status":     room.Status,
		}).
		Where("id=?", room.ID)
}

func roomUpdateStateQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at": room.UpdatedAt,
			"state":      room.State,
		}).
		Where("id=?", room.ID)
}

func roomUpdateBookingsQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at":  room.UpdatedAt,
			"booking_ids": room.Bookings,
		}).
		Where("id=?", room.ID)
}

func roomDeleteBookingsQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		Set("booking_ids", nil).
		Where("id=?", room.ID)
}

func roomUpdateLinksQuery(room dbmodel.Room) sq.UpdateBuilder {
	return base.Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at": room.UpdatedAt,
			"links":      room.Links,
		}).
		Where("id=?", room.ID)
}

func roomUpdateParticipantsQuery(room dbmodel.Room, participant []participantDB) sq.UpdateBuilder {
	return base.
		Builder().
		Update(roomTableName).
		SetMap(map[string]interface{}{
			"updated_at":   room.UpdatedAt,
			"participants": participant,
		}).
		Where("id=?", room.ID)
}
