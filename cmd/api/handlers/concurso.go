package handlers

import (
	"encoding/json"
	"net/http"

	"go.concurco_vaga.railway/cmd/api/scrapers"
)

func GetConcursosByLocalidade(w http.ResponseWriter, r *http.Request) {
	localidade := r.URL.Query().Get("localidade")
	if localidade == "" {
		http.Error(w, "Parâmetro 'localidade' é obrigatório", http.StatusBadRequest)
		return
	}

	concursos, err := scrapers.ScrapeConcursosByLocalidade(localidade)
	if err != nil {
		http.Error(w, "Erro ao buscar concursos", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(concursos)
}

func GetConcursosByCategoria(w http.ResponseWriter, r *http.Request) {
	localidade := r.URL.Query().Get("localidade")
	area := r.URL.Query().Get("area")
	if localidade == "" || area == "" {
		http.Error(w, "Parâmetros 'localidade' e 'area' são obrigatórios", http.StatusBadRequest)
		return
	}

	concursos, err := scrapers.ScrapeConcursosByCategoriaAndLocalidade(localidade, area)
	if err != nil {
		http.Error(w, "Erro ao buscar concursos", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(concursos)
}

func GetCargosInConcursos(w http.ResponseWriter, r *http.Request) {
	concursos, err := scrapers.ScrapeCargosInConcursos()
	if err != nil {
		http.Error(w, "Erro ao buscar cargos em concursos", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(concursos)
}
