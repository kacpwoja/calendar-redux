package handlers

import (
	"strconv"
	"net/http"
	"encoding/json"
	"time"
	"github.com/beevik/guid"
	
    "github.com/kacpwoja/calendar-redux/server/models"
	"github.com/kacpwoja/calendar-redux/server/eventbase"

	"log"
)

func GetBusyDays(w http.ResponseWriter, r *http.Request) {
	log.Print("busy")
	vals := r.URL.Query()

	log.Print(vals)

	year_s := vals.Get("year")
	month_s := vals.Get("month")
	if year_s == "" || month_s == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	year, err_y := strconv.Atoi(year_s)
	month, err_m := strconv.Atoi(month_s)
	if err_y != nil || err_m != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	busy_days, err := eventbase.GetEventsMonth(year, month)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print(busy_days)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(busy_days)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	log.Print("events")
	vals := r.URL.Query()

	year_s := vals.Get("year")
	month_s := vals.Get("month")
	day_s := vals.Get("day")
	if year_s == "" || month_s == "" || day_s == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	year, err_y := strconv.Atoi(year_s)
	month, err_m := strconv.Atoi(month_s)
	day, err_d := strconv.Atoi(day_s)
	if err_y != nil || err_m != nil || err_d != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	events, err := eventbase.GetEventsDay(date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log.Print("POST")
	vals := r.URL.Query()

	year_s := vals.Get("year")
	month_s := vals.Get("month")
	day_s := vals.Get("day")
	if year_s == "" || month_s == "" || day_s == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("POST BR1")
		return
	}

	year, err_y := strconv.Atoi(year_s)
	month, err_m := strconv.Atoi(month_s)
	day, err_d := strconv.Atoi(day_s)
	if err_y != nil || err_m != nil || err_d != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("POST BR2")
		return
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	log.Printf("%d %d %d", year, month, day)
	log.Print("Date" + date)

	log.Print(r.Body)
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("POST BR3")
		log.Print(err)
		return
	}

	event.ID = guid.New().String()

	err = eventbase.InsertEvent(event.ID, date, event.Time, event.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("POST ISE")
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	log.Print("PUT")
	vals := r.URL.Query()

	year_s := vals.Get("year")
	month_s := vals.Get("month")
	day_s := vals.Get("day")
	if year_s == "" || month_s == "" || day_s == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err_y := strconv.Atoi(year_s)
	_, err_m := strconv.Atoi(month_s)
	_, err_d := strconv.Atoi(day_s)
	if err_y != nil || err_m != nil || err_d != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = eventbase.UpdateEvent(event.ID, event.Time, event.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveEvent(w http.ResponseWriter, r *http.Request) {
	log.Print("DELETE")
	vals := r.URL.Query()

	year_s := vals.Get("year")
	month_s := vals.Get("month")
	day_s := vals.Get("day")
	id := vals.Get("id")
	if year_s == "" || month_s == "" || day_s == "" || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err_y := strconv.Atoi(year_s)
	_, err_m := strconv.Atoi(month_s)
	_, err_d := strconv.Atoi(day_s)

	if err_y != nil || err_m != nil || err_d != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := eventbase.DeleteEvent(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}