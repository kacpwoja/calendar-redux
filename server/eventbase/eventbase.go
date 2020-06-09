package eventbase

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
	"github.com/kacpwoja/calendar-redux/server/models"
	
	"log"
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
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
    err = db.Ping()
    if err != nil {
        return nil, err
	}
	log.Print("Connected to database")
	return db, nil
}

func InsertEvent(id string, date string, time_at string, name string) error {
	ev_date, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	ev_time, err := time.Parse("15:04:05", time_at)
	if err != nil {
		return err
	}

	sqlStatement := `
	INSERT INTO events (ev_id, ev_date, ev_time, ev_name)
	VALUES ($1, $2, $3, $4);`
	_, err = db.Exec(sqlStatement, id, ev_date, ev_time, name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvent(id string, time_at string, name string) error {
	ev_time, err := time.Parse("15:04:05", time_at)
	if err != nil {
		return err
	}

	sqlStatement := `
	UPDATE events
	SET ev_time = $2, ev_name = $3
	WHERE ev_id = $1;`
	_, err = db.Exec(sqlStatement, id, ev_time, name)
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

	ev_date, err := time.Parse("2006-01-02", date)
	if err != nil {
		return events, nil
	}
	rows, err := db.Query("SELECT ev_id, ev_date, ev_time, ev_name FROM events")
	if err != nil {
		return events, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var time_at, date_at time.Time
		err = rows.Scan(&id, &date_at, &time_at, &name)
		if err != nil {
			return events, err
		}

		y1, m1, d1 := date_at.Date()
		y2, m2, d2 := ev_date.Date()

		if y1 == y2 && m1 == m2 && d1 == d2 {
			ev := models.Event{
				ID: id,
				Time: time_at.Format("15:04:05"),
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

func GetEventsMonth(year int, month int) ([]int, error) {
	busy_days := make([]int, 0)

	rows, err := db.Query("SELECT ev_id, ev_date, ev_time, ev_name FROM events")
	if err != nil {
		return busy_days, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var time_at, date_at time.Time
		err = rows.Scan(&id, &date_at, &time_at, &name)
		if err != nil {
			return busy_days, err
		}
		if date_at.Year() == year && date_at.Month() == time.Month(month) {
			day := date_at.Day()
			busy_days = append(busy_days, day)
		}
	}
	err = rows.Err()
	if err != nil {
		return busy_days, err
	}

	return busy_days, nil
}