package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thezeeshan/api/internal/api"
	"github.com/thezeeshan/api/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplcation() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	workoutHandler := api.NewWorkoutHanlder()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
