# 🔗 Go URL Shortener

A minimal, production-ready **URL shortening service** written in **Go**, using **Redis** as the backend storage. It supports generating short URLs and redirecting them back to their original long URLs similar to services like Bit.ly or TinyURL.

## ✨ Features

- Generate short codes for any URL
- Redirect users to the original URL
- URL expiry support (24h default)
- Built using clean architecture (Handler → Service → Repository)
- Fast and lightweight — powered by Go and Redis

---

## 📁 Project Structure

```
.
├── cmd/             # Entry point of the application
│   └── main.go
├── handler/         # HTTP handlers (request/response)
│   └── handler.go
├── repository/      # Redis data persistence
│   └── redis.go
├── service/         # Business logic
│   └── url\_service.go
├── utils/           # Utility functions (hashing, encoding)
│   └── hash.go
├── docker/          # Dockerfile and Docker Compose setup
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env             # Environment variables
├── go.mod           # Go modules
└── README.md

````

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/JawherKl/system-design-study.git
cd system-design-study/implements/systems/go/url-shortener
````

### 2. Configure environment

Create a `.env` file:

```env
PORT=8080
POSTGRES_DSN=postgres://user:password@localhost:5432/url_shortener?sslmode=disable
BASE_URL=http://localhost:8080
REDIS_ADDR=redis:6379
```

### 3. Run with Docker

```bash
docker-compose up --build
```

> This will spin up both the Go app and Redis using Docker.

---

## 📫 API Endpoints

### `POST /shorten`

**Request**

```json
{
  "url": "https://example.com/my-long-url"
}
```

**Response**

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

### `GET /{code}`

**Redirects** the short code back to the original long URL.

---

## 🧪 Example Usage (cURL)

```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url": "https://openai.com"}'
```

---

## 🛠 Tech Stack

* **Language**: Go
* **Database**: Redis
* **Deployment**: Docker, Docker Compose

---

## 🧼 To-Do / Improvements

* [ ] Add support for custom short codes
* [ ] Add click statistics tracking
* [+] Add persistent storage fallback (PostgreSQL)
* [ ] Add tests (unit + integration)

---

## 📄 License

This project is licensed under the MIT License.
