package repository

import (
	"context"
	"database/sql"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"example.com/go-inventory-grpc/ent"
	"github.com/pkg/errors"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Key string

const (
	Transaction Key = "tx"
)

func NewPostgres(ctx context.Context, databaseURI string) (*ent.Client, error) {
	client, err := open(databaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection")
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed creating schema resources")

	}
	return client, nil
}

func open(databaseURL string) (entClient *ent.Client, err error) {
	var db *sql.DB
	db, err = sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup database connection")
	}

	for i := 1; i < 4; i++ {
		if err = db.Ping(); err != nil {
			if i == 3 {
				return nil, errors.Wrap(err, "failed to ping database")
			}
			<-time.After(time.Duration(int32(i*2)) + time.Second)
			continue
		}
		break
	}

	//Create an ent.Driver from 'db'.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil

}

//WithTx includes a database transcation on context
func WithTx(ctx context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(ctx, Transaction, tx)
}

//GetTx gets the database transaction off context
func GetTx(ctx context.Context) (*ent.Tx, error) {
	tx, ok := ctx.Value(Transaction).(*ent.Tx)
	if !ok {
		return nil, errors.New("failed to assert *ent.Tx from context value")
	}
	return tx, nil
}
