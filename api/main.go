package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/thezeeshan/api/internal/app"
	"github.com/thezeeshan/api/internal/routes"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplcation()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Printf("we are running out app %d", port)

	r := routes.SetUpRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Microsecond,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}
