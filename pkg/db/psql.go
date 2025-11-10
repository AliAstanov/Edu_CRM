package db

import (
	"context"
	"fmt"

	"github.com/AliAstanov/Edu_CRM/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnToDb(pgCfg config.PgConfig) (*pgxpool.Pool, error){
	ctx := context.Background()

	url:= fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)
	pool, err := pgxpool.New(ctx,url)
	if err != nil{
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Successfully connected to PostgreSQL")
	return pool, nil
}
