package dbmodel

import (
	"github.com/google/uuid"

	"be/internal/model"
)

type (
	Participant struct {
		UserID uuid.UUID `db:"user_id"`
		User   *User     `db:"user"`
		Role   string    `db:"role"`
	}

	ParticipantList []Participant
)

func (p ParticipantList) Participants() []model.Participant {
	participants := make([]model.Participant, 0, len(p))

	for _, participant := range p {
		participants = append(participants, model.Participant{
			User: participant.User.ToModel(),
			Role: participant.Role,
		})
	}

	return participants
}

func ParticipantsFromModel(participants []model.Participant) ParticipantList {
	participantList := make(ParticipantList, 0, len(participants))

	for _, participant := range participants {
		participantList = append(participantList, Participant{
			UserID: participant.User.ID,
			User:   UserFromModelPtr(participant.User),
			Role:   participant.Role,
		})
	}

	return participantList
}
