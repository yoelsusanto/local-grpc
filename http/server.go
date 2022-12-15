package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yoelsusanto/local-grpc/work"
)

type server struct{}

type SimpleRequest struct {
	Text string `json:"text"`
}

type SimpleResponse struct {
	Text string `json:"text"`
}

func (s *server) HandleSimple(w http.ResponseWriter, req *http.Request) {
	sr := SimpleRequest{}

	err := json.NewDecoder(req.Body).Decode(&sr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(SimpleResponse{
		Text: work.SimpleProcedure(sr.Text),
	}); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (s *server) HandleComplex(w http.ResponseWriter, req *http.Request) {
	sr := SimpleRequest{}

	err := json.NewDecoder(req.Body).Decode(&sr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(SimpleResponse{
		Text: work.HashProcedure(sr.Text),
	}); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func StartHTTPServer(port int) {
	s := &server{}

	mux := http.NewServeMux()
	mux.HandleFunc("/simple", s.HandleSimple)

	log.Printf("Listening at port %d", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux)
	if err != nil {
		log.Fatalf("failed to listen at port: %d", port)
	}
}
