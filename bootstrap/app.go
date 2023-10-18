package bootstrap

import "github.com/dilyara4949/clean-project/mongo"

type Application struct {
	Env *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application)CloseDBConnection() {
	app.CloseDBConnection()
}