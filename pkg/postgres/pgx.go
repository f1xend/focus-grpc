package postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type PgConn struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func NewPg(ctx context.Context, conn PgConn, l *log.Logger) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.Db)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		err = fmt.Errorf("error connect to DB (%s): %w", conn.Host, err)
		l.Println(err)
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		err = fmt.Errorf("error connect to DB (%s): %w", conn.Host, err)
		l.Println(err)
		return nil, err
	}
	l.Println("connection to DB is ok", conn.Host, conn.Db)
	return db, nil
}
