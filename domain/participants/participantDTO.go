package participants

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

type Participant struct {
	ID                int64  `json:"id"`
	ParticipantNumber string `json:"participant_number"`
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	CreatedAt         string `json:"created_at"`
}

func (participant *Participant) Validate() *errors.RestErr {
	participant.ParticipantNumber = strings.TrimSpace(participant.ParticipantNumber)
	participant.Name = strings.TrimSpace(participant.Name)

	if participant.ParticipantNumber == "" {
		return errors.NewBadRequestError("invalid participant number")
	}
	if participant.Name == "" {
		return errors.NewBadRequestError("invalid name")
	}
	return nil
}