package main

import (
	"api/internal/configuration"
	dbUtil "api/internal/db"
	"api/internal/db/migration"
	"api/internal/endpoints/bakingschedule"
	"api/internal/endpoints/dayoff"
	"api/internal/endpoints/order"
	"api/internal/endpoints/pastry"
	auth "api/internal/middlewares"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	configuration.LoadConfiguration()

	if _, err := os.Stat(configuration.AppConfig.Database.Path); err != nil {
		db = dbUtil.InitDb()
	} else {
		db = dbUtil.ConnectToDb()
	}
	if err := migration.RunMigrations("migration", db); err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	pastry.DB = db
	order.DB = db
	bakingschedule.DB = db
	dayoff.DB = db
	dayoff.DB = db

	router := mux.NewRouter()
	router.HandleFunc("/pastry", pastry.GetPastries).Methods("GET")
	router.HandleFunc("/pastry", pastry.UpdatePastry).Methods("PUT")
	router.HandleFunc("/order", order.GetOrders).Methods("GET")
	router.HandleFunc("/order", order.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", order.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/order/schedule", order.ScheduleOrder).Methods("POST")
	router.HandleFunc("/schedule", bakingschedule.GetBakingSchedules).Methods("GET")
	router.HandleFunc("/schedule", bakingschedule.CreateBakingSchedules).Methods("POST")
	router.HandleFunc("/schedule", bakingschedule.UpdateBakingSchedule).Methods("PUT")
	router.HandleFunc("/dayoff", dayoff.GetDayOffs).Methods("GET")
	router.HandleFunc("/dayoff", dayoff.CreateDayOff).Methods("POST")
	router.HandleFunc("/dayoff/{id}", dayoff.DeleteDayOff).Methods("DELETE")

	authMiddleware := auth.NewAuth(router,
		configuration.AppConfig.Auth.Token,
		configuration.AppConfig.Auth.AdminToken,
		configuration.AppConfig.Auth.NonAdminEndpoints)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"AuthToken"},
		AllowCredentials: true,
	}).Handler(authMiddleware)

	log.Println("Listening on :5555")
	http.ListenAndServe(":5555", corsMiddleware)
}
