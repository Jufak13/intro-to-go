package handler

import (
	"encoding/json"
	"intro-to-go/internal/model"
	"net/http"
)

type Handler struct {
	todos map[int]string
}

func (h *Handler) GetTodos(writer http.ResponseWriter, request *http.Request) {
	h.todos = map[int]string{
		1: "zbud se",
		2: "pojdi kakat",
		3: "Ojoj! zajtrk",
	}

	response := &model.TodosResponse{}
	// Check if todos are nil or empty
	if len(h.todos) == 0 {
		response.Message = "no todos in memory"
		writer.WriteHeader(http.StatusNotFound)
		
	
	} else {
		response.Todos = h.todos
	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	} 

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)
	}
	}
