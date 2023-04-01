package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"example.com/go-inventory-grpc/config"
	"example.com/go-inventory-grpc/ent"
	"github.com/pkg/errors"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Key string

const (
	Transaction Key = "tx"
)

type DB interface {
	GetEntClient() *ent.Client
	IsAlive() bool
	NewTransaction(ctx context.Context) (*ent.Tx, error)
}

type db struct {
	uri         string
	db          *sql.DB
	entClient   *ent.Client
	alive       bool
	schemaAdded bool

	mu *sync.Mutex
}

func NewPostgres(ctx context.Context, databaseURI string) (DB, error) {
	db := db{uri: databaseURI, mu: &sync.Mutex{}}
	defer func() {
		//watch established connection in go routine
		db.watchConnection()
		db.connectionPlan(ctx)
	}()

	entClient, sqlDB, err := db.open(ctx)
	if err != nil {
		return &db, errors.Wrap(err, "failed opening connection")
	}

	if err := entClient.Schema.Create(ctx); err != nil {
		return &db, errors.Wrap(err, "failed creating schema resources")

	}

	db.alive = true
	db.db = sqlDB
	db.entClient = entClient

	return &db, nil
}

func (d *db) watchConnection() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			if err := d.pingDB(); err != nil {
				fmt.Println("Failed to ping database")
				d.alive = false

			} else {
				d.alive = true
			}
		}
	}()
}

func (d *db) connectionPlan(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for range ticker.C {
			if !d.IsAlive() {
				entClient, db, err := d.open(ctx)
				if err != nil {
					d.alive = false
					fmt.Println("Failed to establish DB connection")

				} else {
					d.mu.Lock()
					d.closeEnt()
					d.entClient = entClient
					d.db = db
					d.alive = true
					d.mu.Unlock()
				}
			}
		}
	}()
}

func (d *db) IsAlive() bool {
	if !d.schemaAdded {
		return false
	}
	return d.alive
}

func (d *db) pingDB() error {
	if d.db == nil {
		return errors.New("failed to find db to ping")
	}

	return d.db.Ping()
}

func (d *db) closeEnt() {
	if d.entClient != nil {
		d.entClient.Close()
	}
}

func (d *db) open(ctx context.Context) (entClient *ent.Client, db *sql.DB, err error) {
	if !d.IsAlive() {
		db, err = sql.Open("pgx", d.uri)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to setup connection")
		}

		if err = db.PingContext(ctx); err != nil {
			return nil, db, err
		}

		//creating an ent.Driver from db
		drv := entsql.OpenDB(dialect.Postgres, db)

		entClient := ent.NewClient(ent.Driver(drv))
		if !d.schemaAdded {
			if err := entClient.Schema.Create(ctx); err != nil {
				return nil, nil, errors.New("failed to create schema")
			}
			d.schemaAdded = true
		}

		return ent.NewClient(ent.Driver(drv)), db, nil

	}

	return nil, nil, errors.New("connection already open")
}

func (d *db) validate() error {
	if d.db == nil {
		return errors.New("failed to find db connection")
	}
	if d.entClient == nil {
		return errors.New("failed to find initialized entClient")
	}

	if !d.IsAlive() {
		return errors.New("failed tp find established connection")
	}

	return nil
}

func (d *db) GetEntClient() *ent.Client {
	return d.entClient
}

func (d *db) NewTransaction(ctx context.Context) (*ent.Tx, error) {

	if err := d.validate(); err != nil {
		return nil, err
	}
	return d.entClient.Tx(ctx)
}

// WithTx includes a database transcation on context
func WithTx(ctx context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(ctx, Transaction, tx)
}

// GetTx gets the database transaction off context
func GetTx(ctx context.Context) (*ent.Tx, error) {
	tx, ok := ctx.Value(Transaction).(*ent.Tx)
	if !ok {
		return nil, errors.New("failed to assert *ent.Tx from context value")
	}
	return tx, nil
}

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewEntClient(cfg *config.Config) (DB, error) {
	db := db{}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBConnection.Host, cfg.DBConnection.Port, cfg.DBConnection.User, cfg.DBConnection.Password, cfg.DBConnection.Dbname)

	client, err := ent.Open("postgres", psqlInfo, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	db.alive = true
	//db.db = sqlDB
	db.entClient = client

	return &db, nil
}
