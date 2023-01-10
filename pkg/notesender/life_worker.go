package notesender

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"be/internal/bot"
	"be/internal/datastore/base"
	"be/internal/model"
)

func (s *NoteSender) lifeWorker(ctx context.Context) {
	ticker := time.NewTicker(s.cfg.LifePeriod)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.lifeSend(ctx)
			}
		}
	}()
}

//nolint:gocognit
func (s *NoteSender) lifeSend(ctx context.Context) {
	notes, err := s.store.Address().LifeNotesList(ctx)
	if err != nil {
		s.logger.Err(err).Msg("error get life for notify")
	}

	for _, note := range notes {
		note.Status = model.Sent

		if note.ChannelID == uuid.Nil {
			channelID, err := s.store.Address().GetChannelFromLifeNote(ctx, &note.RecipientID)
			if err != nil && !errors.Is(err, base.ErrNotFound) {
				s.logger.Error().Err(err).Msg("error get channel for life note")
				continue
			}

			if channelID == uuid.Nil {
				channelIDStr, err := s.makeStream(ctx, note.RecipientID)
				if err != nil {
					s.logger.Error().Err(err).Msg("error make stream for life note")
					continue
				}

				channelID = uuid.MustParse(channelIDStr)
			}

			note.ChannelID = channelID
		}

		if err := s.sendNoteToLife(note); err != nil {
			s.logger.Error().Err(err).Msg("error send note to life")

			note.Status = model.SendFailed
		}

		if _, err := s.store.Address().UpdateLifeNote(ctx, *note); err != nil {
			s.logger.Err(err).Msgf("error update email note, id %s, to status %s", note.ID, note.Status)
		}
	}
}

func (s *NoteSender) sendNoteToLife(n *model.LifeNote) error {
	teardown, err := s.login()
	if err != nil {
		return err
	}

	defer teardown()

	message, entities := bot.BuildMessageEntities(n.Body, n.ForEntities)

	if _, err = s.chatSender.PublishStream(&bot.CreateMessageRequest{
		Content:  message,
		Entities: entities,
		StreamID: n.ChannelID.String(),
		Type:     1,
	}); err != nil {
		return err
	}

	return nil
}

func (s *NoteSender) login() (tearDownFunc func(), err error) {
	req := &bot.LoginBotRequest{
		ID: s.cfg.BotID,
		IP: s.cfg.IP,
		OS: bot.OS{
			Name:    runtime.GOOS,
			Version: runtime.GOARCH,
		},
		Secret: s.cfg.BotToken,
	}

	if _, err = s.chatSender.Login(req); err != nil {
		s.logger.Err(err).Msg("login")
		return nil, err
	}

	return func() {
		s.chatSender.Logout()
	}, nil
}

//nolint:gocognit
func (s *NoteSender) makeStream(ctx context.Context, userID uuid.UUID) (string, error) {
	user, err := s.store.User().FindUserByPK(ctx, userID)
	if err != nil {
		return "", err
	}

	if user.ProfileID == "" {
		s.logger.Info().Msg("profile id empty, skip unregistered user...")
		return "", errors.New("profile id empty")
	}

	teardown, err := s.login()
	if err != nil {
		return "", err
	}

	defer teardown()

	streamID := ""

	res, err := s.chatSender.GetSubscribedStreams()
	if err != nil {
		return "", err
	}

	req := &bot.CreateStreamRequest{
		Description: fmt.Sprintf("Проектный офис (уведомления для %s)", user.ProfileID),
		Invites:     []string{user.ProfileID},
		Title:       "Проектный офис",
		Type:        Channel,
		Visible:     Private,
	}

	for _, stream := range res.Streams {
		if stream.Stream.Description == req.Description && stream.Stream.Title == req.Title {
			streamID = stream.Stream.ID
			break
		}
	}

	if streamID == "" {
		resp, err := s.chatSender.CreateStream(req)
		if err != nil {
			return "", err
		}

		streamID = resp.ID
	}

	return streamID, nil
}
