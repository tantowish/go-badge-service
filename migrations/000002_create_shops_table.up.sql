CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    badge_id INT REFERENCES badges(id) ON DELETE SET NULL
);