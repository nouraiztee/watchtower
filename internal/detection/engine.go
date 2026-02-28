package detection

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Engine struct {
	db *pgxpool.Pool
}

func NewEngine(db *pgxpool.Pool) *Engine {
	return &Engine{db: db}
}

func (e *Engine) Start() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for range ticker.C {
			e.checkBruteForce()
		}
	}()
}

func (e *Engine) checkBruteForce() {
	query := `
	SELECT user_id, COUNT(*) 
	FROM events
	WHERE event_type = 'login_attempt'
	AND status = 'failed'
	AND timestamp >= NOW() - INTERVAL '5 minutes'
	GROUP BY user_id
	HAVING COUNT(*) >= 5;
	`

	rows, err := e.db.Query(context.Background(), query)
	if err != nil {
		log.Println("Detection query failed:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		var count int

		err := rows.Scan(&userID, &count)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		log.Printf("🚨 ALERT: Possible brute-force attack on user %s (%d failed attempts)\n", userID, count)
	}
}