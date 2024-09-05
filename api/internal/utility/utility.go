package utility

import (
	"log"
	"net/http"
)

func GetLanguageOrDefault(request *http.Request) string {
	languageCode := request.Header.Get("Accept-Language")
	if languageCode == "" {
		languageCode = "en"
	}
	return languageCode
}

func LogAndErrorResponse(err error, response http.ResponseWriter) {
	LogAndErrorResponseWithCode(err, response, http.StatusInternalServerError)
}

func LogAndErrorResponseWithCode(err error, response http.ResponseWriter, responseCode int) {
	log.Println(err.Error())
	http.Error(response, "", responseCode)
}
