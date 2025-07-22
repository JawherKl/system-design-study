# System Design Study with Multi-Language Implementations

![Repository Size](https://img.shields.io/github/repo-size/JawherKl/system-design-study)
![Last Commit](https://img.shields.io/github/last-commit/JawherKl/system-design-study)
![Issues](https://img.shields.io/github/issues-raw/JawherKl/system-design-study)
![Forks](https://img.shields.io/github/forks/JawherKl/system-design-study)
![Stars](https://img.shields.io/github/stars/JawherKl/system-design-study)

![system-design-study](https://github.com/JawherKl/system-design-study/blob/main/images/system-design.png)

Welcome to the **System Design Study** project — a comprehensive collection of real-world system design challenges, explored and implemented across multiple programming languages: **PHP (Symfony)**, **Go**, **JavaScript (Node.js)**, and **Java (Spring Boot)**.

This repository offers **practical examples, theoretical insights, architecture diagrams**, and **multi-tech implementations** to help you master scalable, reliable, and maintainable system design.

---

## Table of Contents

1. [Introduction](#introduction)  
2. [Project Structure](#project-structure)  
3. [System Design Categories & Cases](#system-design-categories--cases)  
4. [Supported Languages](#supported-languages)  
5. [How to Use](#how-to-use)  
6. [Contributing](#contributing)  
7. [License](#license)

---

## Introduction

System design is a fundamental skill for software engineers and architects, involving the creation of scalable, fault-tolerant, and maintainable software systems.  
This project dives into diverse design cases, combining **theory, diagrams, and code implementations** in four languages to strengthen your understanding from multiple angles.

Each case includes:

- **Theoretical explanations and system diagrams**  
- **Detailed real-world problem statements**  
- **Clean, tested code implementations**  
- **Multi-language support with best practices**

---

## Project Structure

```plaintext
system-design-study/
├── README.md                      # Project overview and instructions
├── diagrams/                      # Architecture and system design diagrams
├── docs/                          # Theoretical explanations and use cases
├── implements/
│   ├── platforms/                # Foundational building blocks
│   | ├── project-name/              # Project implementation
│   | |  ├── go/                        # Go implementations
│   | |  ├── nodejs/                    # JavaScript (Node.js) implementations
│   | |  ├── springboot/                # Java (Spring Boot) implementations
│   | |  └── symfony/                   # PHP (Symfony) implementations
│   ├── services/                 # Focused microservices
│   | ├── project-name/              # Project implementation
│   | |  ├── go/                        # Go implementations
│   | |  ├── nodejs/                    # JavaScript (Node.js) implementations
│   | |  ├── springboot/                # Java (Spring Boot) implementations
│   | |  └── symfony/                   # PHP (Symfony) implementations
│   ├── system/                   # Full end-to-end applications
│   | ├── project-name/              # Project implementation
│   | |  ├── go/                        # Go implementations
│   | |  ├── nodejs/                    # JavaScript (Node.js) implementations
│   | |  ├── springboot/                # Java (Spring Boot) implementations
─   ─ ─  └── symfony/                   # PHP (Symfony) implementations
```

---

## System Design Categories & Cases

We classify system design challenges into **three main categories** to better understand their scope and purpose:

---

### 🧩 1. Systems
> **Full end-to-end applications** solving real-world problems. These systems include multiple components and services.

| Idea | Description | Priority |
|------|-------------|----------|
| `url-shortener` | Create a scalable short URL system like Bitly | 1 |
| `real-time-chat` | Real-time messaging system with WebSocket & delivery tracking | 3 |
| `ecommerce-backend` | Modular backend for cart, order, and payment systems | 4 |
| `file-storage-system` | Upload, store, and manage user files securely (like Dropbox) | 5 |
| `notification-system` | Send emails, SMS, and push messages via async queues | 2 |
| `blogging-system` | Markdown-based blog system with publishing workflows | 12 |
| `document-collab` | Real-time collaborative editing with sync and conflict resolution | 15 |
| `news-feed` | Build a social media system with feeds, content scoring, large-scale reads | 7 |

---

### 🏗️ 2. Platforms
> **Foundational building blocks** that power other systems. Reusable, scalable, and often infrastructure-level.

| Idea | Description | Priority |
|------|-------------|----------|
| `multi-tenant-saas` | Design a SaaS-ready backend with tenant isolation | 11 |
| `analytics-engine` | Track user events & pageviews with scalable storage | 14 |
| `video-streaming` | Stream videos with encoding, metadata, and CDN-like delivery | 8 |
| `online-booking` | Prevent double booking with concurrency control (tickets/hotels) | 9 |
| `online-exam` | Secure online exam platform with timed sessions and grading | 20 |

---

### ⚙️ 3. Services
**Focused microservices** providing specific capabilities, often embedded inside larger systems.

| Idea | Description | Priority |
|------|-------------|----------|
| `rate-limiter` | Throttle API usage per user/IP using sliding window or token bucket | 6 |
| `auth-service` | Centralized authentication & authorization (JWT, OAuth2) | 13 |
| `job-queue` | Queue and process background jobs with retries and scheduling | 10 |
| `payment-service` | Handle secure payments with idempotency and webhook confirmation | 19 |
| `search-service` | Full-text search with ranking and autocomplete (Elastic-like) | 16 |
| `iot-ingestion` | Collect high-throughput telemetry from IoT devices | 18 |
| `cdn-file-serving` | Deliver static/media files with edge caching & TTL | 17 |

---

## Supported Languages

| Language                 | Folder Name   | Build Command                     | Test Command          |
| ------------------------ | ------------- | --------------------------------- | --------------------- |
| **PHP (Symfony)**        | `symfony/`    | `php bin/console`                 | `php bin/phpunit`     |
| **Go**                   | `go/`         | `go build ./src/...`              | `go test ./tests/...` |
| **JavaScript (Node.js)** | `nodejs/`     | `npm run build` or `node src/...` | `npm test`            |
| **Java (Spring Boot)**   | `springboot/` | `./mvnw clean install`            | `./mvnw test`         |

---

## How to Use

### 1. Clone the Repository

```bash
git clone https://github.com/JawherKl/system-design-study.git
cd system-design-study
```

### 2. Choose a System Design Case

Navigate to the implementation in your preferred language, for example:

```bash
cd implements/nodejs/url-shortener
```

### 3. Follow the Instructions

Each implementation folder contains detailed setup, build, and run instructions, including:

* Environment variables
* Dependencies
* Running servers and tests

### 4. Explore & Learn

* Review the theory and diagrams in `/docs` and `/diagrams`
* Compare implementations across languages
* Experiment and extend the systems!

---

## Contributing

Contributions are highly encouraged! To contribute:

1. Fork the repository
2. Create a feature branch (`feature/your-feature-name`)
3. Implement your additions or fixes
4. Submit a pull request with clear descriptions

---

## License

This project is licensed under the [MIT License](./LICENSE).

---

## Contact

Feel free to open issues or pull requests for questions, suggestions, or collaboration.
