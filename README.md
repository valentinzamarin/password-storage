# Password Storage

This is a small application written in Golang that I created for personal use and with the desire to write something in this language for myself.

Stack: 
- Golang
- SQLite
- GORM
- Fyne.io
- DDD

```go build -ldflags="-H=windowsgui" -o password-storage.exe .\cmd\app\main.go```

---

Added an Event Bus to avoid updating list in each component separately