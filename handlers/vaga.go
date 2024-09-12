package handlers

import (
	"concurso_vaga/scrapers"
	"encoding/json"
	"net/http"
)

func GetVagas(w http.ResponseWriter, r *http.Request) {
	vagas, err := scrapers.ScrapeVagas()
	if err != nil {
		http.Error(w, "Erro ao buscar vagas", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(vagas)
}
