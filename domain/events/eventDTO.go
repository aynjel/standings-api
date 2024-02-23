package events

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

// Event

type Event struct {
	ID          int64  `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	OrganizerID int64  `json:"organizer_id"`
	CreatedAt   string `json:"created_at"`
}

func (event *Event) Validate() *errors.RestErr {
	event.Title = strings.TrimSpace(event.Title)
	event.Description = strings.TrimSpace(event.Description)
	event.Location = strings.TrimSpace(event.Location)
	event.StartDate = strings.TrimSpace(event.StartDate)
	event.EndDate = strings.TrimSpace(event.EndDate)

	if event.Title == "" {
		return errors.NewBadRequestError("invalid title")
	}
	if event.Description == "" {
		return errors.NewBadRequestError("invalid description")
	}
	if event.Location == "" {
		return errors.NewBadRequestError("invalid location")
	}
	if event.StartDate == "" {
		return errors.NewBadRequestError("invalid start date")
	}
	if event.EndDate == "" {
		return errors.NewBadRequestError("invalid end date")
	}
	return nil
}

// EventOrganizer

type EventOrganizer struct {
	ID        int64  `json:"ID"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

func (eventOrganizer *EventOrganizer) Validate() *errors.RestErr {
	eventOrganizer.Name = strings.TrimSpace(eventOrganizer.Name)
	eventOrganizer.Email = strings.TrimSpace(eventOrganizer.Email)
	eventOrganizer.Phone = strings.TrimSpace(eventOrganizer.Phone)

	if eventOrganizer.Name == "" {
		return errors.NewBadRequestError("invalid name")
	}
	if eventOrganizer.Email == "" {
		return errors.NewBadRequestError("invalid email")
	}
	if eventOrganizer.Phone == "" {
		return errors.NewBadRequestError("invalid phone")
	}
	return nil
}

// EventParticipant

type EventParticipant struct {
	ID            int64  `json:"ID"`
	EventID       int64  `json:"event_id"`
	ParticipantID int64  `json:"participant_id"`
	CreatedAt     string `json:"created_at"`
}

func (eventParticipant *EventParticipant) Validate() *errors.RestErr {
	if eventParticipant.EventID == 0 {
		return errors.NewBadRequestError("invalid event id")
	}
	if eventParticipant.ParticipantID == 0 {
		return errors.NewBadRequestError("invalid participant id")
	}
	return nil
}