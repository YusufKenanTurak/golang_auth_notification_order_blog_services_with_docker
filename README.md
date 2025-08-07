# Go Enterprise Microservices Suite

This repository contains a suite of independent, production-ready Go microservices built with clean architecture and containerization best practices.

## 🧩 Microservices

### 🔐 Auth Service
Basic user authentication (login/register)

Endpoints:
- `POST /login`
- `POST /register`

---

### 📦 Order Service
Simulated order management

Endpoints:
- `GET /orders`

---

### 🔔 Notification Service
Basic notification trigger system

Endpoint:
- `GET /notify`

---

### 📝 Blog Service
Exposes static blog content

Endpoint:
- `GET /posts`

## 🧪 Tech Stack

- Go 1.20+
- Gorilla Mux
- Docker & Docker Compose
- Clean folder structure
- Extensible architecture

## ▶️ How to Run

```bash
docker-compose up --build
Services will be available on:

Auth: http://localhost:8083

Notification: http://localhost:8084

Order: http://localhost:8085

Blog: http://localhost:8086

🚀 Suggested Enhancements
Add gRPC or REST gateway

Add JWT-based authentication

Connect shared PostgreSQL DBs for multi-service state

✨ Author
Yusuf Kenan Turak
Senior Fullstack Developer | Go + .NET | React