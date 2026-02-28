# WatchTower

**Learning SIEM Concepts by Building a Mini Tool (AI-assisted)**

WatchTower is a small project I built to understand **how SIEM (Security Information and Event Management) systems work**.  
It collects login events and detects suspicious activity like repeated failed login attempts.

> Built with **AI guidance**, the goal was to learn SIEM concepts, not to create a production-ready system.

---

## What It Does

- Collects login events via REST API (`/api/logs`)
- Detects repeated failed login attempts
- Prints alerts in real time for suspicious activity
- Helps understand **SIEM concepts**: log collection, pattern detection, alerting

> The backend is written in **Go** with a PostgreSQL database, but the tech is secondary — the focus is **on learning security monitoring**.

---

## Quick Start

**1. Clone the repository:**

```bash
git clone https://github.com/nouraiztee/watchtower.git
cd watchtower
```

**2. Set up environment variables:**

Copy the example file and fill in dummy/test values:

```bash
cp env.example .env
```

Example `.env`:

```env
DATABASE_URL=postgres://username:password@localhost:5432/watchtower
API_KEY=YOUR_API_KEY_HERE
```

**3. Start PostgreSQL (optional, for local testing):**

You can run PostgreSQL in Docker:

```bash
docker run --name watchtower-db \
  -e POSTGRES_USER=watchtower \
  -e POSTGRES_PASSWORD=watchtower \
  -e POSTGRES_DB=watchtower \
  -p 5432:5432 -d postgres
```

**4. Run the server:**

```bash
go run cmd/server/main.go
```

Server will start at `http://localhost:8080`.

---

## Testing WatchTower

Send a test login event with `curl`:

```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_API_KEY_HERE" \
  -d '{
    "source": "test-service",
    "event_type": "login_attempt",
    "user_id": "testuser",
    "ip_address": "192.168.1.1",
    "status": "failed"
  }'
```

Send multiple failed login events quickly — the detection engine will print:

```
🚨 ALERT: Possible brute-force attack on user testuser (5 failed attempts)
```

---

## Project Structure

```
watchtower/
├── cmd/server/main.go       # Entry point
├── internal/config/         # Config loader
├── internal/models/         # Event structs
├── internal/storage/        # DB repository
├── internal/detection/      # Detection logic
├── env.example              # Example environment variables
├── go.mod
├── go.sum
└── README.md
```

---

## Notes

- This project is for **learning SIEM concepts only**
- Built with **AI guidance** to structure the backend
- Replace `.env` placeholders with dummy/test data
- Focus is on understanding security monitoring, not production-ready tooling
