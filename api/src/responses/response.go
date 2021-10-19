package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCodes int, dados interface{}) {
	w.WriteHeader(statusCodes)
	w.Header().Set("Content-Type", "application/json")

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
