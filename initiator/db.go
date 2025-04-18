package initiator

import (
	"context"
	"fmt"
	"helphub-backend/platform/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(url string, log logger.Logger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}

	config.MaxConns = 1000 // Not tested yet
	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}
	return conn
}
