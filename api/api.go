package api

import (
	"encoding/json"
	"net/http"
)

type CardCountParams struct {
	Username string
}

type CardCountResponse struct {
	Code  int
	Count int8
}

type Error struct {
	Code    int
	Message string
}

type Card int

const (
	Hearts Card = iota
	Spades
	Clubs
	Diamonds
)

func (c Card) String() string {
	switch c {
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	default:
		return ""
	}
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
