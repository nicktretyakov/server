package report

import (
	"context"

	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s *Storage) BulkDelete(ctx context.Context, reports []model.ReportBooking) error {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		for _, report := range reports {
			q, err := s.db.ExecBuilder(ctx, deleteQuery(report.ID, report.BookingID, report.Period.Time()))
			if err != nil {
				return err
			}

			if q.RowsAffected() == 0 {
				return base.ErrNotFound
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
