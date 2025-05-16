package boostrap

import "blog/pgsql"

type Aplictation struct {
	Env  *Env
	Psql pgsql.Client
}

func App() *Aplictation {
	app := &Aplictation{}
	app.Env = NewEnv()
	app.Psql = NewPsql(app.Env)
	return app
}

func (app *Aplictation) CloseBDConnection() {
	ClosePsqlConnection(app.Psql)
}
