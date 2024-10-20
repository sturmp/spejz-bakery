package main

import (
	"api/internal/configuration"
	dbUtil "api/internal/db"
	"api/internal/db/migration"
	"api/internal/endpoints/bakingschedule"
	"api/internal/endpoints/dayoff"
	"api/internal/endpoints/order"
	"api/internal/endpoints/pastry"
	"api/internal/endpoints/unitofmeasure"
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

	order.DB = db
	bakingschedule.DB = db
	unitofmeasure.DB = db

	router := mux.NewRouter()
	router.HandleFunc("/order", order.GetOrders).Methods("GET")
	router.HandleFunc("/order", order.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", order.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/order/complete/{id}", order.CompleteOrder).Methods("PUT")
	router.HandleFunc("/order/schedule", order.ScheduleOrder).Methods("POST")
	router.HandleFunc("/schedule", bakingschedule.GetBakingSchedules).Methods("GET")
	router.HandleFunc("/schedule", bakingschedule.CreateBakingSchedules).Methods("POST")
	router.HandleFunc("/schedule", bakingschedule.UpdateBakingSchedule).Methods("PUT")
	router.HandleFunc("/schedule", bakingschedule.DeleteBakingSchedule).Methods("DELETE")
	router.HandleFunc("/unitofmeasure", unitofmeasure.GetUnitOfMeasures).Methods("GET")

	pastry.RegisterHandler(router, db)
	dayoff.RegisterHandler(router, db)

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
