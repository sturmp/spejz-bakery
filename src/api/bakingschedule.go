package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type BakingSchedule struct {
	Pastry    string
	Quantity  float32
	Reserved  float32
	ReadyDate time.Time
}

func GetBakingSchedules(response http.ResponseWriter, request *http.Request) {
	rows, err := db.Query("select pastry, quantity, reserved, readydate from bakingschedule")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bakingschedules := []BakingSchedule{}
	for rows.Next() {
		var bakingschedule BakingSchedule
		var readyDateText string
		err = rows.Scan(&bakingschedule.Pastry, &bakingschedule.Quantity, &bakingschedule.Reserved, &readyDateText)
		if err != nil {
			log.Fatal(err)
		}
		if readyDate, err := time.Parse(time.RFC3339, readyDateText); err == nil {
			bakingschedule.ReadyDate = readyDate
		} else {
			log.Println(err)
		}
		bakingschedules = append(bakingschedules, bakingschedule)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(bakingschedules)
}
