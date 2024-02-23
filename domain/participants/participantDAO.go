package participants

import "anggi.tabulation/database/postgreSQL/tabulationDB"

var (
	createTableParticipant = `
		CREATE TABLE IF NOT EXISTS participants (
			id SERIAL PRIMARY KEY,
			participant_number VARCHAR(100) NOT NULL,
			name VARCHAR(100) NOT NULL,
			phone VARCHAR(20) NULLABLE,
			email VARCHAR(100) NULLABLE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	queryInsertParticipant  = `INSERT INTO participants (participant_number, name, phone, email) VALUES ($1, $2, $3, $4) RETURNING id`
	queryGetParticipantByID = `SELECT id, participant_number, name, phone, email, created_at FROM participants WHERE id = $1`
	queryGetAllParticipants = `SELECT id, participant_number, name, phone, email, created_at FROM participants`
	queryUpdateParticipant  = `UPDATE participants SET participant_number = $1, name = $2, phone = $3, email = $4 WHERE id = $5`
	queryDeleteParticipant  = `DELETE FROM participants WHERE id = $1`
)

func init() {
	if _, err := tabulationDB.Client.Exec(createTableParticipant); err != nil {
		panic(err)
	}
}

func (participant *Participant) Save() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(participant.ParticipantNumber, participant.Name, participant.Phone, participant.Email).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (participant *Participant) GetByID() (*Participant, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetParticipantByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Participant
	if err := stmt.QueryRow(participant.ID).Scan(&p.ID, &p.ParticipantNumber, &p.Name, &p.Phone, &p.Email, &p.CreatedAt); err != nil {
		return nil, err
	}

	return &p, nil
}

func GetAll() ([]Participant, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllParticipants)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []Participant
	for rows.Next() {
		var p Participant
		if err := rows.Scan(&p.ID, &p.ParticipantNumber, &p.Name, &p.Phone, &p.Email, &p.CreatedAt); err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

func (participant *Participant) Update() error {
	stmt, err := tabulationDB.Client.Prepare(queryUpdateParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(participant.ParticipantNumber, participant.Name, participant.Phone, participant.Email, participant.ID); err != nil {
		return err
	}

	return nil
}

func (participant *Participant) Delete() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(participant.ID); err != nil {
		return err
	}

	return nil
}