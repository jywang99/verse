package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"jy.org/verse/src/config"
	"jy.org/verse/src/logging"
)

var cfg = config.Config.Db
var logger = logging.Logger

type DbConn struct {
    conn *pgx.Conn
}

var Conn = &DbConn{}
func Setup() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
    url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode, cfg.Schema)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
        logger.ERROR.Fatalf("Error connecting to database: %v\n", url)
        logger.ERROR.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
    logger.INFO.Printf("Connected to database: %v\n", url)
    Conn.conn = conn
}

func (db *DbConn) Close() {
    db.conn.Close(context.Background())
}

