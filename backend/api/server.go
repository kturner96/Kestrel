package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer(serverPort string, pool *pgxpool.Pool) {


	http.ListenAndServe(serverPort, nil)
}