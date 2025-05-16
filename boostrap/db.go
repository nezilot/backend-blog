package boostrap

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func NewPsql(env *Env) *sql.DB {

	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName
	dbSsl := env.DbSslMode

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPass, dbName, dbSsl)

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("ошибка при подключении к psql: %w", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("ошибка при проверке подключения к psql: %w", err))
	}

	return db
}
