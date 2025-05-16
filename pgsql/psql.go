package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
	Collection(string) Collection
}

type Collection interface {
	FindOne(context.Context, interface{}) (SingleResult, error)
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	UpdateOne(context.Context, interface{}, interface{}) (int64, error)
}

type SingleResult interface {
	Scan(...interface{}) error
}

type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Disconnect(context.Context) error
}

type postgresClient struct {
	db *sql.DB
}

type postgresDatabase struct {
	db *sql.DB
}

type postgresCollection struct {
	db *sql.DB
}

func NewClient(connection string) (Client, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к PostgreSQL: %w", err)
	}
	return &postgresClient{db: db}, nil
}

func (pc *postgresClient) Database(dbName string) Database {
	return &postgresDatabase{db: pc.db}
}

func (pc *postgresClient) Connect(ctx context.Context) error {
	return pc.db.PingContext(ctx)
}

func (pc *postgresClient) Disconnect(ctx context.Context) error {
	return pc.db.Close()
}

func (pd *postgresDatabase) Collection(colName string) Collection {
	return &postgresCollection{db: pd.db}
}

func (pc *postgresCollection) FindOne(ctx context.Context, filter interface{}) (SingleResult, error) {
	query := "SELECT * FROM your_table WHERE your_condition"
	row := pc.db.QueryRowContext(ctx, query)
	return &postgresSingleResult{row: row}, nil
}

func (pc *postgresCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	query := "INSERT INTO your_table (columns) VALUES ($1) RETURNING id"
	var id int
	err := pc.db.QueryRowContext(ctx, query, document).Scan(&id)
	return id, err
}

func (pc *postgresCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	query := "DELETE FROM your_table WHERE your_condition"
	res, err := pc.db.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (pc *postgresCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (int64, error) {
	query := "UPDATE your_table SET column = $1 WHERE your_condition"
	res, err := pc.db.ExecContext(ctx, query, update)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

type postgresSingleResult struct {
	row *sql.Row
}

func (sr *postgresSingleResult) Scan(dest ...interface{}) error {
	return sr.row.Scan(dest...)
}
