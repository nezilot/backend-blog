package boostrap

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Aplictation struct {
	Env  *Env
	Psql sql.DB
}

func App() Aplictation {
	app := &Aplictation{}
	app.Env = NewEnv()
	app.Psql = *NewPsql(app.Env)
	return *app
}
