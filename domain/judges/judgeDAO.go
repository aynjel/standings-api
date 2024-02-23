package judges

import "anggi.tabulation/database/postgreSQL/tabulationDB"

var (
	createTableJudge = `CREATE TABLE IF NOT EXISTS judges (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		phone VARCHAR(20) NOT NULL,
		email VARCHAR(100) NOT NULL,
		username VARCHAR(100) NOT NULL,
		pin INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	createTableJudgeEvent = `CREATE TABLE IF NOT EXISTS judge_events (
		id SERIAL PRIMARY KEY,
		event_id INT NOT NULL,
		judge_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (judge_id) REFERENCES judges(id)
	);`
	queryInsertJudge  = `INSERT INTO judges (name, phone, email, username, pin) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	queryGetJudgeByID = `SELECT id, name, phone, email, username, pin, created_at FROM judges WHERE id = $1`
	queryGetAllJudges = `SELECT id, name, phone, email, username, pin, created_at FROM judges`
	queryUpdateJudge  = `UPDATE judges SET name = $1, phone = $2, email = $3, username = $4, pin = $5 WHERE id = $6`
	queryDeleteJudge  = `DELETE FROM judges WHERE id = $1`

	queryInsertJudgeEvent        = `INSERT INTO judge_events (event_id, judge_id) VALUES ($1, $2) RETURNING id`
	queryGetJudgeEvents          = `SELECT id, event_id, judge_id, created_at FROM judge_events WHERE event_id = $1`
	queryGetJudgeEventsByJudgeID = `SELECT id, event_id, judge_id, created_at FROM judge_events WHERE judge_id = $1`
	queryDeleteJudgeEvent        = `DELETE FROM judge_events WHERE event_id = $1 AND judge_id = $2`
)

func init() {
	if _, err := tabulationDB.Client.Exec(createTableJudge); err != nil {
		panic(err)
	}
	if _, err := tabulationDB.Client.Exec(createTableJudgeEvent); err != nil {
		panic(err)
	}
}

func (judge *Judge) SaveJudge() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertJudge)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(judge.Name, judge.Phone, judge.Email, judge.Username, judge.Pin).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (judge *Judge) GetByID() (*Judge, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetJudgeByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var j Judge
	if err := stmt.QueryRow(judge.ID).Scan(&j.ID, &j.Name, &j.Phone, &j.Email, &j.Username, &j.Pin, &j.CreatedAt); err != nil {
		return nil, err
	}

	return &j, nil
}

func GetAll() ([]Judge, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllJudges)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var judges []Judge
	for rows.Next() {
		var j Judge
		if err := rows.Scan(&j.ID, &j.Name, &j.Phone, &j.Email, &j.Username, &j.Pin, &j.CreatedAt); err != nil {
			return nil, err
		}
		judges = append(judges, j)
	}

	return judges, nil
}

func (judge *Judge) UpdateJudge() error {
	stmt, err := tabulationDB.Client.Prepare(queryUpdateJudge)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(judge.Name, judge.Phone, judge.Email, judge.Username, judge.Pin, judge.ID); err != nil {
		return err
	}

	return nil
}

func (judge *Judge) DeleteJudge() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteJudge)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(judge.ID); err != nil {
		return err
	}

	return nil
}

func (judgeEvent *JudgeEvent) SaveJudgeEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryInsertJudgeEvent)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(judgeEvent.EventID, judgeEvent.JudgeID); err != nil {
		return err
	}

	return nil
}

func (judgeEvent *JudgeEvent) GetJudgeEvents() ([]JudgeEvent, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetJudgeEvents)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(judgeEvent.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var judgeEvents []JudgeEvent
	for rows.Next() {
		var je JudgeEvent
		if err := rows.Scan(&je.ID, &je.EventID, &je.JudgeID, &je.CreatedAt); err != nil {
			return nil, err
		}
		judgeEvents = append(judgeEvents, je)
	}

	return judgeEvents, nil
}

func (judgeEvent *JudgeEvent) GetJudgeEventsByJudgeID() ([]JudgeEvent, error) {
	stmt, err := tabulationDB.Client.Prepare(queryGetJudgeEventsByJudgeID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(judgeEvent.JudgeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var judgeEvents []JudgeEvent
	for rows.Next() {
		var je JudgeEvent
		if err := rows.Scan(&je.ID, &je.EventID, &je.JudgeID, &je.CreatedAt); err != nil {
			return nil, err
		}
		judgeEvents = append(judgeEvents, je)
	}

	return judgeEvents, nil
}

func (judgeEvent *JudgeEvent) DeleteJudgeEvent() error {
	stmt, err := tabulationDB.Client.Prepare(queryDeleteJudgeEvent)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(judgeEvent.EventID, judgeEvent.JudgeID); err != nil {
		return err
	}

	return nil
}