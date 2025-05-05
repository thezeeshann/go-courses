package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/thezeeshan/api/internal/app"
)

func SetUpRoutes(app *app.Application) *chi.Mux {

	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	return r
}
