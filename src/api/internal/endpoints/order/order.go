package order

import (
	"api/internal/configuration"
	"api/internal/endpoints/bakingschedule"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wneessen/go-mail"
)

var DB *sql.DB

type Order struct {
	Pastry        string
	Customer      string
	Quantity      float32
	PreferedDate  time.Time
	ScheduledDate time.Time
}

func GetOrders(response http.ResponseWriter, request *http.Request) {
	orders := fetchOrdersFromDB()

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(orders)
}

func CreateOrder(response http.ResponseWriter, request *http.Request) {
	var order Order

	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	schedules := bakingschedule.FetchSchedulesFromDB()
	for _, schedule := range schedules {
		if IsOrderFitInSchedule(order, schedule) {
			order.ScheduledDate = order.PreferedDate
			schedule.Reserved += order.Quantity
			bakingschedule.UpdateScheduleReservedInDB(schedule)
		}
	}
	InsertOrderToDb(order)

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(order)

	go sendEmail(order)
}

func fetchOrdersFromDB() []Order {
	rows, err := DB.Query("select pastry, customer, quantity, preferedDate, scheduledDate from pastryorder")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		var preferedDateText string
		var scheduledDateText string
		err = rows.Scan(&order.Pastry, &order.Customer, &order.Quantity, &preferedDateText, &scheduledDateText)
		if err != nil {
			log.Fatal(err)
		}

		if preferedDate, err := time.Parse(time.RFC3339, preferedDateText); err == nil {
			order.PreferedDate = preferedDate
		} else {
			log.Println(err)
		}

		if scheduledDate, err := time.Parse(time.RFC3339, scheduledDateText); err == nil {
			order.ScheduledDate = scheduledDate
		} else {
			order.ScheduledDate = time.Time{}
		}

		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return orders
}

func IsOrderFitInSchedule(order Order, schedule bakingschedule.BakingSchedule) bool {
	return schedule.Pastry == order.Pastry &&
		schedule.ReadyDate.UTC() == order.PreferedDate.UTC() &&
		schedule.Quantity-schedule.Reserved >= order.Quantity
}

func InsertOrderToDb(order Order) {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`insert into
		pastryorder(pastry, customer, quantity, preferedDate, scheduledDate)
        values(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.Pastry, order.Customer, order.Quantity, order.PreferedDate.Format(time.RFC3339), order.ScheduledDate.Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func sendEmail(order Order) {
	config := configuration.AppConfig

	email := mail.NewMsg()
	if err := email.From(config.Email.From); err != nil {
		log.Printf("Failed to set From: %s", err)
	}
	if err := email.To(config.Email.To); err != nil {
		log.Printf("Failed to set To: %s", err)
	}
	email.Subject("New Bakery Order from " + order.Customer)
	email.SetBodyString(mail.TypeTextPlain, fmt.Sprintf("%s\n%s\n%f\n%s", order.Customer, order.Pastry, order.Quantity, order.PreferedDate.Format("2006-01-02 15:04")))

	client, err := mail.NewClient(config.Email.Smtp.Host, mail.WithPort(config.Email.Smtp.Port), mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(config.Email.Smtp.User), mail.WithPassword(config.Email.Smtp.Pass))
	if err != nil {
		log.Printf("failed to create mail client: %s", err)
	}

	if err := client.DialAndSend(email); err != nil {
		log.Printf("failed to send mail: %s", err)
	}
}
