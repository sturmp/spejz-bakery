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
	log.Println(err.Error())
	http.Error(response, "", http.StatusInternalServerError)
}
