package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"fds/src/controllers"
	"fds/src/repositories"
)

func main() {
	controller := controllers.New(&repositories.UserRepository{})

	db, connError := sql.Open("mysql", "root:root@/go")

	if connError != nil {
		panic(connError)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		err := controller.Store(w, r, db)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})

	r.Get("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		user, err := controller.GetByName(w, r, db, chi.URLParam(r, "name"))

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		userJson, err := json.Marshal(user)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		w.Write(userJson)
	})

	http.ListenAndServe(":3000", r)
}
