package organizersauth

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

type Organizer struct {
	ID               int64  `json:"id"`
	EventOrganizerID int64  `json:"event_organizer_id"`
	Username         string `json:"username"`
	Pin         int    `json:"pin"`
	CreatedAt        string `json:"created_at"`
}

func (organizer *Organizer) Validate() *errors.RestErr {
	organizer.Username = strings.TrimSpace(organizer.Username)

	if organizer.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}
	if organizer.Pin == 0 {
		return errors.NewBadRequestError("invalid pin")
	}
	
	return nil
}