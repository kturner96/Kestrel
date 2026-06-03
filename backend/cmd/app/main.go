package main

import (
	"github.com/kturner96/kestrel/backend/api"
	"github.com/kturner96/kestrel/backend/cfg"
	"github.com/kturner96/kestrel/backend/internal/db"
)

func main()  {
	config := cfg.LoadCfg()
	pool := db.OpenDb(config.DbConnection)
	defer pool.Close()
	api.StartServer(config.Port, pool)

}