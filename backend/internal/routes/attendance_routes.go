package routes

import (
	"test-webdev-suiten-25/internal/controller"

	"github.com/gorilla/mux"
)

func RegisterAttendanceRoutes(router *mux.Router, controller controller.AttendanceController) {
	router.HandleFunc("/api/attendance", controller.GetByDateAndDivision).Methods("GET")
	router.HandleFunc("/api/attendance/bulk", controller.BulkUpsert).Methods("POST")
}
