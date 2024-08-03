package order

import (
	"api/internal/configuration"
	"api/internal/endpoints/bakingschedule"
	"api/internal/endpoints/pastry"
	"api/internal/utility"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wneessen/go-mail"
)

var DB *sql.DB

type Order struct {
	Id     int
	Pastry struct {
		Id   int
		Name string
	}
	Customer      string
	Quantity      float32
	PreferedDate  time.Time
	ScheduledDate time.Time
}

type CreateOrderRequest struct {
	PastryId     int
	Customer     string
	Quantity     float32
	PreferedDate time.Time
}

type ScheduleOrderRequest struct {
	Id            int
	ScheduledDate time.Time
}

func GetOrders(response http.ResponseWriter, request *http.Request) {
	orders, err := fetchOrdersFromDB()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(orders)
}

func CreateOrder(response http.ResponseWriter, request *http.Request) {
	var order CreateOrderRequest

	if err := json.NewDecoder(request.Body).Decode(&order); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	schedules, err := bakingschedule.FetchSchedulesFromDB()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	var scheduledDate time.Time
	for _, schedule := range schedules {
		if isOrderFitInSchedule(order.PastryId, order.Quantity, schedule, order.PreferedDate) {
			schedule.Reserved += order.Quantity
			scheduledDate = schedule.ReadyDate
			if err := bakingschedule.UpdateScheduleReservedInDB(schedule); err != nil {
				utility.LogAndErrorResponse(err, response)
			}
			break
		}
	}

	if err := insertOrderToDb(order, scheduledDate); err != nil {
		utility.LogAndErrorResponse(err, response)
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(order)

	go sendEmail(order)
}

func ScheduleOrder(response http.ResponseWriter, request *http.Request) {
	var scheduleOrderRequest ScheduleOrderRequest
	if err := json.NewDecoder(request.Body).Decode(&scheduleOrderRequest); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := fetchOrdersFromDB()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	var order Order
	for i := 0; i < len(orders); i++ {
		if orders[i].Id == scheduleOrderRequest.Id {
			order = orders[i]
			break
		}
	}

	schedules, err := bakingschedule.FetchSchedulesFromDB()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
	}
	for _, schedule := range schedules {
		if isOrderFitInSchedule(order.Pastry.Id, order.Quantity, schedule, scheduleOrderRequest.ScheduledDate) {
			order.ScheduledDate = scheduleOrderRequest.ScheduledDate
			schedule.Reserved += order.Quantity
			if err := bakingschedule.UpdateScheduleReservedInDB(schedule); err != nil {
				utility.LogAndErrorResponse(err, response)
			}
			if err := updateOrderScheduleDateInDB(order); err != nil {
				utility.LogAndErrorResponse(err, response)
			}
			break
		}
	}

	encoder := json.NewEncoder(response)
	encoder.SetIndent("", "  ")
	encoder.Encode(order)
}

func DeleteOrder(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	orderId, ok := vars["id"]
	if !ok {
		http.Error(response, "Missing Order id!", http.StatusBadRequest)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	stmt, err := tx.Prepare(`DELETE from pastryorder WHERE id = ?`)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderId)
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}

	err = tx.Commit()
	if err != nil {
		utility.LogAndErrorResponse(err, response)
		return
	}
}

func fetchOrdersFromDB() ([]Order, error) {
	rows, err := DB.Query(`SELECT
			pastryorder.id,
			pastryorder.pastryid,
			pastry.name,
			pastryorder.customer,
			pastryorder.quantity,
			pastryorder.preferedDate,
			pastryorder.scheduledDate FROM pastryorder
		JOIN pastry ON pastryorder.pastryid = pastry.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		var preferedDateText string
		var scheduledDateText string
		err = rows.Scan(&order.Id,
			&order.Pastry.Id,
			&order.Pastry.Name,
			&order.Customer,
			&order.Quantity,
			&preferedDateText,
			&scheduledDateText)
		if err != nil {
			return nil, err
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
		return nil, err
	}
	return orders, nil
}

func updateOrderScheduleDateInDB(order Order) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE pastryorder SET scheduledDate = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ScheduledDate.Format(time.RFC3339), order.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func isOrderFitInSchedule(pastryId int,
	quantity float32,
	schedule bakingschedule.BakingSchedule,
	scheduleDate time.Time) bool {
	return schedule.Pastry.Id == pastryId &&
		schedule.ReadyDate.UTC() == scheduleDate.UTC() &&
		schedule.Quantity-schedule.Reserved >= quantity
}

func insertOrderToDb(order CreateOrderRequest, scheduledDate time.Time) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO
		pastryorder(pastryid, customer, quantity, preferedDate, scheduledDate)
        VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	scheduledDateString := ""
	if !scheduledDate.IsZero() {
		scheduledDateString = scheduledDate.Format((time.RFC3339))
	}
	_, err = stmt.Exec(order.PastryId, order.Customer, order.Quantity, order.PreferedDate.Format(time.RFC3339), scheduledDateString)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func sendEmail(order CreateOrderRequest) {
	config := configuration.AppConfig
	pastryName := pastry.FetchPastryName(order.PastryId)

	email := mail.NewMsg()
	if err := email.From(config.Email.From); err != nil {
		log.Printf("Failed to set From: %s", err)
	}
	if err := email.To(config.Email.To); err != nil {
		log.Printf("Failed to set To: %s", err)
	}
	email.Subject("New Bakery Order from " + order.Customer)
	email.SetBodyString(mail.TypeTextPlain, fmt.Sprintf("%s\n%s\n%f\n%s", order.Customer, pastryName, order.Quantity, order.PreferedDate.Format("2006-01-02 15:04")))

	client, err := mail.NewClient(config.Email.Smtp.Host, mail.WithPort(config.Email.Smtp.Port), mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(config.Email.Smtp.User), mail.WithPassword(config.Email.Smtp.Pass))
	if err != nil {
		log.Printf("failed to create mail client: %s", err)
	}

	if err := client.DialAndSend(email); err != nil {
		log.Printf("failed to send mail: %s", err)
	}
}
