package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/models"
	"github.com/gorilla/mux"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := models.Todos{
		models.Todo{Name: "Write presentation"},
		models.Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
