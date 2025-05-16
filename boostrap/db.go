package boostrap

import (
	"blog/pgsql"
	"context"
	"fmt"
	"log"
	"time"
)

func NewPsql(env *Env) pgsql.Client {
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName
	dbSsl := env.DbSslMode

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPass, dbName, dbSsl)

	client, err := pgsql.NewClient(connStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func ClosePsqlConnection(client pgsql.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

}
