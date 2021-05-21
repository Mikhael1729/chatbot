package handlers

import (
	"github.com/Mikhael1729/restaurant-chatbot/models"
	"log"
	"net/http"
)

type Messages struct {
	logger *log.Logger
}

func NewMessages(logger *log.Logger) *Messages {
	return &Messages{logger}
}

func (handler *Messages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	setupCors(rw, r)
	if r.Method == http.MethodGet {
		handler.getMessages(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		handler.addMessage(rw, r)
		return
	}

	// Catch all.
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getMessages process the GET method for the handler
func (m *Messages) getMessages(rw http.ResponseWriter, r *http.Request) {
	messages := models.GetMessages()
	err := messages.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal messages json", http.StatusInternalServerError)
	}
}

// getMessages process the POST method for the handler
func (handler *Messages) addMessage(rw http.ResponseWriter, r *http.Request) {
	message := &models.Message{}

	err := message.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	models.AddMessage(message)
}

func setupCors(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}