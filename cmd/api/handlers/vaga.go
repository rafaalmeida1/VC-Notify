package handlers

import (
	"encoding/json"
	"net/http"

	"go.concurco_vaga.railway/cmd/api/scrapers"
)

func GetVagas(w http.ResponseWriter, r *http.Request) {
	vagas, err := scrapers.ScrapeVagas()
	if err != nil {
		http.Error(w, "Erro ao buscar vagas", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(vagas)
}
