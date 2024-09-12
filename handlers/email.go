package handlers

import (
	"concurso_vaga/utils"
	"encoding/json"
	"net/http"
)

type NotificationRequest struct {
	Email         string        `json:"email"`
	Tipo          string        `json:"tipo"`
	Oportunidades []Opportunity `json:"oportunidades"`
}

type Opportunity struct {
	Titulo string `json:"titulo"`
	Link   string `json:"link"`
}

func SendNotification(w http.ResponseWriter, r *http.Request) {
	var req NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Tipo == "" || len(req.Oportunidades) == 0 {
		http.Error(w, "Campos 'email', 'tipo' e 'oportunidades' são obrigatórios", http.StatusBadRequest)
		return
	}

	mensagem := ""
	for _, op := range req.Oportunidades {
		mensagem += op.Titulo + " - " + op.Link + "\n"
	}
	assunto := "Novas oportunidades de " + req.Tipo

	if err := utils.SendEmail(req.Email, assunto, mensagem); err != nil {
		http.Error(w, "Erro ao enviar e-mail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
