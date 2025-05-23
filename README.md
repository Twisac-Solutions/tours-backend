# 🧭 Tours Backend API

A REST API for managing tours, events, destinations, and reviews — built with **Go**, **Fiber**, **GORM**, and **SQLite**.

## 🚀 Features

- Email & Google Auth (JWT-based)
- CRUD: Users, Tours, Events, Destinations, Categories, Reviews
- Modular structure
- Swagger API docs (`/swagger/index.html`)
- Live reload with `air`

## 📦 Tech Stack

Go, Fiber, GORM, SQLite, Swagger, Air

## ⚙️ Getting Started

```bash
git clone https://github.com/Twisac-Solutions/tours-backend.git
cd tours-backend
go mod tidy
air
````

## 📚 Swagger Docs

```bash
swag init --generalInfo cmd/main.go --output docs
```

View docs at `http://localhost:8000/swagger/index.html`

---

MIT © 2025 [Twisac Solutions](https://github.com/Twisac-Solutions)
