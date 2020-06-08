package eventbase

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
    "github.com/kacpwoja/calendar-redux/server/models"
)

const (
	host		= "db"
	port		= 5432
	user		= "postgres"
	password	= "test"
	dbname		= "postgres"
)

var db *sql.DB

func Init() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
    err = db.Ping()
    if err != nil {
        return nil, err
	}
	return db, nil
}

func InsertEvent(id string, date string, time_at string, name string) error {
	ev_date, _ := time.Parse("2020-06-10", date)
	ev_time, _ := time.Parse("16:20", time_at)

	sqlStatement := `
	INSERT INTO events (ev_id, ev_date, ev_time, ev_name)
	VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(sqlStatement, id, ev_date, ev_time, name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvent(id string, time_at string, name string) error {
	ev_time, _ := time.Parse("16:20", time_at)

	sqlStatement := `
	UPDATE events
	SET ev_time = $2, ev_name = $3
	WHERE ev_id = $1;`
	_, err := db.Exec(sqlStatement, id, ev_time, name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(id string) error {
	sqlStatement := `
	DELETE FROM events
	WHERE ev_id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func GetEventsDay(date string) ([]models.Event, error) {
	events := make([]models.Event, 0)

	ev_date, _ := time.Parse("2020-06-10", date)
	rows, err := db.Query("SELECT ev_id, ev_date, ev_time, ev_name FROM events")
	if err != nil {
		return events, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var time_at, date_at time.Time
		err = rows.Scan()
		if err != nil {
			return events, err
		}
		if date_at == ev_date {
			ev := models.Event{
				ID: id,
				Time: time_at.Format("16:20"),
				Name: name,
			}
			events = append(events, ev)
		}
	}
	err = rows.Err()
	if err != nil {
		return events, err
	}
	return events, nil
}

func GetEventsMonth(year int, month int) ([]models.Event, error) {
	events := make([]models.Event, 0)

	rows, err := db.Query("SELECT ev_id, ev_date, ev_time, ev_name FROM events")
	if err != nil {
		return events, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var time_at, date_at time.Time
		err = rows.Scan()
		if err != nil {
			return events, err
		}
		if date_at.Year() == year && date_at.Month() == time.Month(month) {
			ev := models.Event{
				ID: id,
				Time: time_at.Format("16:20"),
				Name: name,
			}
			events = append(events, ev)
		}
	}
	err = rows.Err()
	if err != nil {
		return events, err
	}
	return events, nil
}