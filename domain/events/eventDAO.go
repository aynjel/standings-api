package events

import "anggi.tabulation/database/postgreSQL/tabulationDB"

var (
	createTableEvent = `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			description TEXT NOT NULL,
			location VARCHAR(100) NOT NULL,
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP NOT NULL,
			organizer_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (organizer_id) REFERENCES event_organizers(id)
		);
	`
	createTableEventOrganizer = `
		CREATE TABLE IF NOT EXISTS event_organizers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
			phone VARCHAR(20) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	createEventParticipantTable = `
		CREATE TABLE IF NOT EXISTS event_participants (
			id SERIAL PRIMARY KEY,
			event_id INT NOT NULL,
			participant_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (event_id) REFERENCES events(id),
			FOREIGN KEY (participant_id) REFERENCES participants(id)
		);
	`
	queryInsertEvent  = `INSERT INTO events (title, description, location, start_date, end_date, organizer_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	queryGetEventByID = `SELECT id, title, description, location, start_date, end_date, organizer_id, created_at FROM events WHERE id = $1`
	queryGetAllEvents = `SELECT id, title, description, location, start_date, end_date, organizer_id, created_at FROM events`
	queryUpdateEvent  = `UPDATE events SET title = $1, description = $2, location = $3, start_date = $4, end_date = $5 WHERE id = $6`
	queryDeleteEvent  = `DELETE FROM events WHERE id = $1`

	queryInsertEventOrganizer  = `INSERT INTO event_organizers (name, email, phone) VALUES ($1, $2, $3) RETURNING id`
	queryGetEventOrganizerByID = `SELECT id, name, email, phone, created_at FROM event_organizers WHERE id = $1`
	queryGetAllEventOrganizers = `SELECT id, name, email, phone, created_at FROM event_organizers`
	queryUpdateEventOrganizer  = `UPDATE event_organizers SET name = $1, email = $2, phone = $3 WHERE id = $4`
	queryDeleteEventOrganizer  = `DELETE FROM event_organizers WHERE id = $1`

	queryInsertEventParticipant = `INSERT INTO event_participants (event_id, participant_id) VALUES ($1, $2) RETURNING id`
	queryGetEventParticipants    = `SELECT id, event_id, participant_id, created_at FROM event_participants WHERE event_id = $1`
	queryGetParticipantEvents    = `SELECT id, event_id, participant_id, created_at FROM event_participants WHERE participant_id = $1`
	queryDeleteEventParticipant = `DELETE FROM event_participants WHERE event_id = $1 AND participant_id = $2`
)

func init() {
	if _, err := tabulationDB.Client.Exec(createTableEvent); err != nil {
		panic(err)
	}
	if _, err := tabulationDB.Client.Exec(createTableEventOrganizer); err != nil {
		panic(err)
	}
	if _, err := tabulationDB.Client.Exec(createEventParticipantTable); err != nil {
		panic(err)
	}
}

// Event

func (event *Event) SaveEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertEvent)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(event.Title, event.Description, event.Location, event.StartDate, event.EndDate, event.OrganizerID).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (event *Event) GetEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryGetEventByID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(event.ID)
	if getErr := result.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.OrganizerID, &event.CreatedAt); getErr != nil {
		return getErr
	}

	return nil
}

func (event *Event) GetAllEvents() ([]Event, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllEvents)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.OrganizerID, &event.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (event *Event) UpdateEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryUpdateEvent)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(event.Title, event.Description, event.Location, event.StartDate, event.EndDate, event.ID); err != nil {
		return err
	}

	return nil
}

func (event *Event) DeleteEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteEvent)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(event.ID); err != nil {
		return err
	}

	return nil
}

// EventOrganizer

func (eventOrganizer *EventOrganizer) SaveEventOrganizer() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertEventOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(eventOrganizer.Name, eventOrganizer.Email, eventOrganizer.Phone).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (eventOrganizer *EventOrganizer) GetEventOrganizer() error {
	stmt, err := tabulationDB.Client.Prepare(queryGetEventOrganizerByID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(eventOrganizer.ID)
	if getErr := result.Scan(&eventOrganizer.ID, &eventOrganizer.Name, &eventOrganizer.Email, &eventOrganizer.Phone, &eventOrganizer.CreatedAt); getErr != nil {
		return getErr
	}

	return nil
}

func (eventOrganizer *EventOrganizer) GetAllEventOrganizers() ([]EventOrganizer, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllEventOrganizers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventOrganizers []EventOrganizer
	for rows.Next() {
		var eventOrganizer EventOrganizer
		if err := rows.Scan(&eventOrganizer.ID, &eventOrganizer.Name, &eventOrganizer.Email, &eventOrganizer.Phone, &eventOrganizer.CreatedAt); err != nil {
			return nil, err
		}
		eventOrganizers = append(eventOrganizers, eventOrganizer)
	}

	return eventOrganizers, nil
}

func (eventOrganizer *EventOrganizer) UpdateEventOrganizer() error {
	stmt, err := tabulationDB.Client.Prepare(queryUpdateEventOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(eventOrganizer.Name, eventOrganizer.Email, eventOrganizer.Phone, eventOrganizer.ID); err != nil {
		return err
	}

	return nil
}

func (eventOrganizer *EventOrganizer) DeleteEventOrganizer() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteEventOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(eventOrganizer.ID); err != nil {
		return err
	}

	return nil
}

// EventParticipant

func (eventParticipant *EventParticipant) SaveEventParticipant() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertEventParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(eventParticipant.EventID, eventParticipant.ParticipantID); err != nil {
		return err
	}

	return nil
}

func (eventParticipant *EventParticipant) GetEventParticipants() ([]EventParticipant, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetEventParticipants)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventParticipant.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParticipants []EventParticipant
	for rows.Next() {
		var ep EventParticipant
		if err := rows.Scan(&ep.ID, &ep.EventID, &ep.ParticipantID, &ep.CreatedAt); err != nil {
			return nil, err
		}
		eventParticipants = append(eventParticipants, ep)
	}

	return eventParticipants, nil
}

func (eventParticipant *EventParticipant) GetParticipantEvents() ([]EventParticipant, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetParticipantEvents)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventParticipant.ParticipantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParticipants []EventParticipant
	for rows.Next() {
		var ep EventParticipant
		if err := rows.Scan(&ep.ID, &ep.EventID, &ep.ParticipantID, &ep.CreatedAt); err != nil {
			return nil, err
		}
		eventParticipants = append(eventParticipants, ep)
	}

	return eventParticipants, nil
}

func (eventParticipant *EventParticipant) DeleteEventParticipant() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteEventParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(eventParticipant.EventID, eventParticipant.ParticipantID); err != nil {
		return err
	}

	return nil
}