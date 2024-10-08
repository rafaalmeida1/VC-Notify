package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.concurco_vaga.railway/cmd/api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/concursos/localidade", handlers.GetConcursosByLocalidade).Methods("GET")
	r.HandleFunc("/concursos/categoria", handlers.GetConcursosByCategoria).Methods("GET")
	r.HandleFunc("/concursos/cargos", handlers.GetCargosInConcursos).Methods("GET")
	r.HandleFunc("/vagas", handlers.GetVagas).Methods("GET")
	r.HandleFunc("/notificar", handlers.SendNotification).Methods("POST")

	log.Println("Servidor rodando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
