package router

import (
	"github.com/gorilla/mux"
    "github.com/kacpwoja/calendar-redux/server/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/BusyDays", handlers.GetBusyDays).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Events", handlers.GetEvents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/Event", handlers.CreateEvent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/Event", handlers.EditEvent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/Event", handlers.RemoveEvent).Methods("DELETE", "OPTIONS")

	return router
}