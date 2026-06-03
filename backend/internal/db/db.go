package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDb(connString string) (*pgxpool.Pool) {
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	return dbpool

}