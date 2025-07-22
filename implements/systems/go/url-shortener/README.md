### Overview: 
A simplified version of [Bitly](https://bitly.com/) where:
* Users can input a long URL and get a short one (like https://sho.rt/abc123)
* When someone visits the short URL, they are redirected to the original long URL
* The service handles high read traffic with low latency

### Architecture Components:
```plaintext
User ───► API Server (Go)
             ├── Redis Cache (short -> long)
             ├── PostgreSQL (persistence)
             └── Hash Generator (e.g., base62, SHA256, UUID)
```

### Project Structure:
```bash
url-shortener/
├── cmd/                  # Entry point (main.go)
├── internal/
│   ├── handler/          # HTTP handlers (create, redirect)
│   ├── service/          # Business logic (generate short URL, etc.)
│   ├── repository/       # DB and cache access
│   ├── model/            # DTOs / data models
│   └── utils/            # Hashing, encoding, validation
├── pkg/                  # Reusable code (optional)
├── config/               # Config management
├── docker/               # Docker setup
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
└── README.md
```

### Tech Stack:
* Go — Web server and business logic
* PostgreSQL — Store mappings (short → long)
* Redis — Cache for fast lookups
* gorilla/mux — HTTP routing
* godotenv — Load .env config
* go-pg or pgx — PostgreSQL ORM/driver
* redigo or go-redis — Redis client
