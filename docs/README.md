## 🗂️ Categories of System Design

We organize system design challenges into **three core categories**:

---

### 🧩 1. Systems
> **Full end-to-end applications** solving real-world problems. These systems include multiple components and services.

| Idea | Description | Priority |
|------|-------------|
| `url-shortener` | Create a scalable short URL system like Bitly | 1
| `real-time-chat` | Real-time messaging system with WebSocket & delivery tracking | 3
| `ecommerce-backend` | Modular backend for cart, order, and payment systems | 4
| `file-storage-system` | Upload, store, and manage user files securely (like Dropbox) | 5
| `notification-system` | Send emails, SMS, and push messages via async queues | 2
| `blogging-system` | Markdown-based blog system with publishing workflows | 12
| `document-collab` | Real-time collaborative editing with sync and conflict resolution | 15
| `news-feed` | Build a social media system with feeds, content scoring, large-scale reads | 7

---

### 🏗️ 2. Platforms
> **Foundational building blocks** that power other systems. Reusable, scalable, and often infrastructure-level.

| Idea | Description | Priority |
|------|-------------|
| `multi-tenant-saas` | Design a SaaS-ready backend with tenant isolation | 11
| `analytics-engine` | Track user events & pageviews with scalable storage | 14
| `video-streaming` | Stream videos with encoding, metadata, and CDN-like delivery | 8
| `online-booking` | Prevent double booking with concurrency control (tickets/hotels) | 9
| `online-exam` | Secure online exam platform with timed sessions and grading | 20

---

### ⚙️ 3. Services
**Focused microservices** providing specific capabilities, often embedded inside larger systems.

| Idea | Description | Priority |
|------|-------------|
| `rate-limiter` | Throttle API usage per user/IP using sliding window or token bucket | 6
| `auth-service` | Centralized authentication & authorization (JWT, OAuth2) | 13
| `job-queue` | Queue and process background jobs with retries and scheduling | 10
| `payment-service` | Handle secure payments with idempotency and webhook confirmation | 19
| `search-service` | Full-text search with ranking and autocomplete (Elastic-like) | 16
| `iot-ingestion` | Collect high-throughput telemetry from IoT devices | 18
| `cdn-file-serving` | Deliver static/media files with edge caching & TTL | 17
