package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jubaer-ad/goapi/api"
	"github.com/jubaer-ad/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCardCount(w http.ResponseWriter, r *http.Request) {
	var params = api.CardCountParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CardDetails
	tokenDetails = (*database).GetUserCards(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CardCountResponse{
		Code:  http.StatusOK,
		Count: int8((*tokenDetails).Cards),
	}

	w.Header().Set("Content-Type", "Application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
