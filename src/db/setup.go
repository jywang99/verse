package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"jy.org/verse/src/config"
	"jy.org/verse/src/logging"
)

var cfg = config.Config.Db
var logger = logging.Logger

type dbConn struct {
    pool *pgxpool.Pool
}
var Conn = &dbConn{}

type queryTracer struct {
    log *log.Logger
}

func (tracer *queryTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
    tracer.log.Println("SQL query: ", "sql", data.SQL, "args", data.Args)
    return ctx
}

func (tracer *queryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}

func Init() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
    url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Schema)

    // pool config
    dbConfig, err := pgxpool.ParseConfig(url)
    if err!=nil {
        log.Fatal("Failed to create a config, error: ", err)
    }
    poolCfg := cfg.Pool
    dbConfig.MaxConns = int32(poolCfg.MaxConns)
    dbConfig.MinConns = int32(poolCfg.MinConns)
    dbConfig.MaxConnLifetime = poolCfg.MaxConnLifetime
    dbConfig.MaxConnIdleTime = poolCfg.MaxConnIdleTime
    dbConfig.HealthCheckPeriod = poolCfg.HealthCheckPeriod
    dbConfig.ConnConfig.ConnectTimeout = poolCfg.ConnTimeout
    dbConfig.ConnConfig.Tracer = &queryTracer{
        log: logger.INFO,
    }
    dbConfig.ConnConfig.RuntimeParams["search_path"] = cfg.Schema

    connPool,err := pgxpool.NewWithConfig(context.Background(), dbConfig)
    if err!=nil {
        log.Fatal("Error while creating connection to the database!!")
    }
    logger.INFO.Println("Successfully connected to the database: ", url)

    Conn.pool = connPool
}

func (dbc *dbConn) Close() {
    dbc.pool.Close()
}

