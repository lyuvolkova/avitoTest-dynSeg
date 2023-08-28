package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Postgres(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse DatabaseDSN: %w", err)
	}

	config.ConnConfig.PreferSimpleProtocol = true
	config.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on", // будет ли обратная косая черта в обычных строковых константах ('...') восприниматься буквально

		"statement_timeout": strconv.FormatInt((5 * time.Minute).Milliseconds(), 10), // максимальное время выполнения оператора
		"lock_timeout":      strconv.FormatInt((5 * time.Minute).Milliseconds(), 10), // максимальную длительность ожидания (в миллисекундах) любым оператором получения блокировки
	}

	config.MaxConns = 5
	config.MinConns = 1
	config.MaxConnIdleTime = 5 * time.Minute

	return pgxpool.ConnectConfig(ctx, config)
}
