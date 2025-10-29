# 🔐 Password Storage

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://golang.org)
[![Fyne GUI](https://img.shields.io/badge/GUI-Fyne.io-1C71D8?logo=go)](https://fyne.io)
[![GORM](https://img.shields.io/badge/ORM-GORM-69C1B8?logo=go)](https://gorm.io)
[![SQLite](https://img.shields.io/badge/Database-SQLite-003B57?logo=sqlite)](https://sqlite.org)

This is a small application written in Golang that I created for personal use and with the desire to write something in this language for myself.

## 🚀 Tech Stack

### Core Dependencies
- **[Go](https://golang.org)** (1.23.0+) - The programming language
- **[Fyne.io](https://fyne.io)** (v2.6.2) - Cross-platform GUI toolkit
- **[GORM](https://gorm.io)** (v1.30.1) - Object-Relational Mapping library
- **[GORM SQLite Driver](https://gorm.io/driver/sqlite)** (v1.6.0) - Database driver for SQLite
- **[SQLite](https://sqlite.org)** - Embedded SQL database engine
- **[golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)** (v0.41.0) - Cryptographic primitives

### Architecture & Design
- **DDD (Domain-Driven Design)** - Architectural approach
- **Clean Architecture** - Separation of concerns

## 📦 Installation & Build

```bash
# Clone the repository
git clone <your-repo-url>
cd password-storage

# Build the application
go build ./cmd/app/main.go

# Or run directly
go run ./cmd/app/main.go
