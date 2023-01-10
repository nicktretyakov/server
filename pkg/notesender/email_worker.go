package notesender

import (
	"context"
	"time"

	"be/internal/model"
)

func (s *NoteSender) emailWorker(ctx context.Context) {
	ticker := time.NewTicker(s.cfg.EmailPeriod)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.emailsSend(ctx)
			}
		}
	}()
}

func (s *NoteSender) emailsSend(ctx context.Context) {
	notes, err := s.store.Address().EmailNotesList(ctx)
	if err != nil {
		s.logger.Err(err).Msg("error get emails for notify")
	}

	for _, note := range notes {
		go func(note *model.EmailNote) {
			note.Status = model.Sent

			if err := s.emailSender.Send(
				note.Subject,
				note.Body,
				note.RecipientEmail,
				note.SenderEmail); err != nil {
				s.logger.Err(err).Msgf("error send email, note id %s", note.ID)
				note.Status = model.SendFailed
			}

			if _, err := s.store.Address().UpdateEmailNote(ctx, *note); err != nil {
				s.logger.Err(err).Msgf("error update email note, id %s, to status %s", note.ID, note.Status)
			}
		}(note)
	}
}
