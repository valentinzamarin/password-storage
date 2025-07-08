CREATE TABLE IF NOT EXISTS passwords (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    description TEXT
);

CREATE INDEX idx_passwords_url ON passwords(url);