package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"untitled13/internal/config"
)

type Database struct {
	db *pgx.Conn
}

func InitDb(dbConfig config.DatabaseConfig) (*Database, error) {
	ctx := context.Background()

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect")
	}
	return &Database{db: conn}, nil
}
