package organizersauth

import "anggi.tabulation/database/postgreSQL/tabulationDB"

var (
	createTableOrganizer = `
		CREATE TABLE IF NOT EXISTS organizers (
			id SERIAL PRIMARY KEY,
			event_organizer_id INT NOT NULL,
			username VARCHAR(100) NOT NULL,
			pin INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (event_organizer_id) REFERENCES event_organizers(id)
		);
	`
	queryInsertOrganizer  = `INSERT INTO organizers (event_organizer_id, username, pin) VALUES ($1, $2, $3) RETURNING id`
	queryGetOrganizerByID = `SELECT id, event_organizer_id, username, pin, created_at FROM organizers WHERE id = $1`
	queryGetAllOrganizers = `SELECT id, event_organizer_id, username, pin, created_at FROM organizers`
	queryUpdateOrganizer  = `UPDATE organizers SET event_organizer_id = $1, username = $2, pin = $3 WHERE id = $4`
	queryDeleteOrganizer  = `DELETE FROM organizers WHERE id = $1`
	loginQueryOrganizer = `SELECT id, event_organizer_id, username, pin FROM organizers WHERE username = $1 AND pin = $2`
)

func init() {
	if _, err := tabulationDB.Client.Exec(createTableOrganizer); err != nil {
		panic(err)
	}
}

func (organizer *Organizer) Save() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(organizer.EventOrganizerID, organizer.Username, organizer.Pin).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (organizer *Organizer) Update() error {
	stmt, err := tabulationDB.Client.Prepare(queryUpdateOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(organizer.EventOrganizerID, organizer.Username, organizer.Pin, organizer.ID); err != nil {
		return err
	}

	return nil
}

func (organizer *Organizer) Delete() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteOrganizer)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(organizer.ID); err != nil {
		return err
	}

	return nil
}

func (organizer *Organizer) GetByID() (*Organizer, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetOrganizerByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var o Organizer
	if err := stmt.QueryRow(organizer.ID).Scan(&o.ID, &o.EventOrganizerID, &o.Username, &o.Pin, &o.CreatedAt); err != nil {
		return nil, err
	}

	return &o, nil
}

func GetAll() ([]Organizer, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllOrganizers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizers []Organizer
	for rows.Next() {
		var o Organizer
		if err := rows.Scan(&o.ID, &o.EventOrganizerID, &o.Username, &o.Pin, &o.CreatedAt); err != nil {
			return nil, err
		}
		organizers = append(organizers, o)
	}

	return organizers, nil
}

func (organizer *Organizer) LoginOrganizer() (*Organizer, error) {
	stmt, err := tabulationDB.Client.Prepare(loginQueryOrganizer)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var o Organizer
	if err := stmt.QueryRow(organizer.Username).Scan(&o.ID, &o.EventOrganizerID, &o.Username, &o.Pin); err != nil {
		return nil, err
	}

	return &o, nil
}