package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kturner96/kestrel/backend/internal/handlers"
)

func StartServer(serverPort string, pool *pgxpool.Pool) {
	h := handlers.Handler{Pool: pool}
	r := mux.NewRouter()
	r.HandleFunc("/sessions", h.HandlePost)

	http.ListenAndServe(serverPort, r)
}