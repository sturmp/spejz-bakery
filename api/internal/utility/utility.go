package utility

import (
	"log"
	"net/http"
)

func LogAndErrorResponse(err error, response http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(response, "", http.StatusInternalServerError)
}
