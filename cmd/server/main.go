package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nouraiztee/watchtower/internal/config"
	"github.com/nouraiztee/watchtower/internal/detection"
	"github.com/nouraiztee/watchtower/internal/models"
	"github.com/nouraiztee/watchtower/internal/storage"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := storage.NewEventRepository(db)

	engine := detection.NewEngine(db)
	engine.Start()

	r := chi.NewRouter()

	r.Post("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != cfg.APIKey {
    		http.Error(w, "Unauthorized", http.StatusUnauthorized)
        	return
		}

		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		if event.Timestamp.IsZero() {
			event.Timestamp = time.Now()
		}

		if err := repo.Insert(r.Context(), &event); err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	log.Println("WatchTower running on :8080")
	http.ListenAndServe(":8080", r)
}