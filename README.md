# System Design Study with Multi-Language Implementations

Welcome to the **System Design Study** project, where we explore and implement various system design concepts using multiple programming languages: **PHP**, **Go**, **JavaScript**, and **Java**. This repository provides practical examples, theoretical explanations, and system diagrams for a comprehensive understanding of real-world system design cases.

## Table of Contents

1. [Introduction](#introduction)  
2. [Project Structure](#project-structure)  
3. [System Design Cases](#system-design-cases)  
4. [Supported Languages](#supported-languages)  
5. [How to Use](#how-to-use)  
6. [Contributing](#contributing)  
7. [License](#license)

---

## Introduction

System design is an essential aspect of software engineering, involving the creation of scalable, reliable, and maintainable systems. This project delves into various design cases, offering theoretical insights alongside practical implementations in multiple languages.

Each system design case includes:

- **Detailed Theorem Explanation**  
- **Diagrams and Architecture**  
- **Real-World Examples**  
- **Code Implementations**  
- **Tests for Verification**

---

## Project Structure

```plaintext
system-design-study/
├── README.md                # Project overview and instructions
├── diagrams/                # System design diagrams
├── docs/                    # Theoretical explanations and use cases
├── implements/ 
│   ├──  php/                     # PHP implementations
│   ├──  go/                      # Go implementations
│   ├──  js/                      # JavaScript implementations
│   ├──  java/                    # Java implementations
└── .gitignore               # Git ignore file
```

---

## System Design Cases

1. **Authentication System**
   - Theory: Token-based auth, OAuth, JWT
   - Diagram: [Authentication Diagram](./diagrams/authentication-system/authentication-diagram.png)
   - Example: Secured login and session management

2. **Caching System**
   - Theory: Cache types (LRU, LFU), Redis
   - Diagram: [Caching Diagram](./diagrams/caching-system/caching-diagram.png)
   - Example: Distributed cache for high-traffic apps

3. **Load Balancing**
   - Theory: Round-robin, least connections
   - Diagram: [Load Balancing Diagram](./diagrams/load-balancing/load-balancing-diagram.png)
   - Example: Distributing requests across microservices

_... More cases coming soon!_

---

## Supported Languages

| Language      | Folder Name | Build Command               | Test Command                   |
|---------------|--------------|------------------------------|---------------------------------|
| **PHP**       | `php/`       | `php src/yourFile.php`       | `phpunit tests/`               |
| **Go**        | `go/`        | `go build ./src/...`         | `go test ./tests/...`          |
| **JavaScript**| `js/`        | `node src/yourFile.js`       | `npm test`                     |
| **Java**      | `java/`      | `javac src/*.java`           | `java -cp src YourTestClass`   |

---

## How to Use

### 1. Clone the Repository

```bash
git clone https://github.com/JawherKl/system-design-study.git
cd system-design-study
```

### 2. Explore a System Design Case

Navigate to any system design case in your preferred language:

```bash
cd php/authentication-system
```

### 3. Run the Code

Follow the build and run instructions for each language in the relevant folder.

### 4. Run Tests

Execute the provided test suite:

```bash
# Example for PHP
phpunit tests/
```

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a new branch (`feature/your-feature-name`)
3. Commit your changes
4. Push to the branch
5. Create a pull request

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

Feel free to reach out with any questions or suggestions via GitHub Issues or Pull Requests!

---

### Screenshots (Optional)
![System Design Example](./diagrams/example-diagram.png)
