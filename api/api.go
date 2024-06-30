package main

import (
	"api/internal/configuration"
	dbUtil "api/internal/db"
	"api/internal/endpoints/bakingschedule"
	"api/internal/endpoints/dayoff"
	"api/internal/endpoints/order"
	"api/internal/endpoints/pastry"
	auth "api/internal/middlewares"
	"database/sql"
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
	defer db.Close()

	pastry.DB = db
	order.DB = db
	bakingschedule.DB = db
	dayoff.DB = db
	dayoff.DB = db

	router := mux.NewRouter()
	router.HandleFunc("/pastry", pastry.GetPastries).Methods("GET")
	router.HandleFunc("/order", order.GetOrders).Methods("GET")
	router.HandleFunc("/order", order.CreateOrder).Methods("POST")
	router.HandleFunc("/schedule", bakingschedule.GetBakingSchedules).Methods("GET")
	router.HandleFunc("/dayoff", dayoff.GetDayOffs).Methods("GET")

	corsMiddleware := cors.Default().Handler(router)
	authMiddleware := auth.NewAuth(corsMiddleware, configuration.AppConfig.Auth.Token)
	http.ListenAndServe(":5555", authMiddleware)
}
