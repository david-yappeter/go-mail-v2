package service

import (
	"net/http"
	"strings"
)

func HandlerSendEmail(w http.ResponseWriter, r *http.Request) {
	emails := r.FormValue("emails")
	subject := r.FormValue("subject")
	body := r.FormValue("body")

	if checkEmptyString(emails) || checkEmptyString(subject) || checkEmptyString(body) {
		http.Error(w, "Invalid Variable", http.StatusBadRequest)
		return
	}

	go func(emails string, subject string, body string) {
		SendEmail(strings.Split(emails, ","), subject, body)
	}(emails, subject, body)
}

func checkEmptyString(s string) bool {
	return strings.Trim(s, " ") == ""
}
