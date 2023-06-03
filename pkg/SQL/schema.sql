CREATE TABLE IF NOT EXISTS tractors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    info jsonb
);

CREATE TABLE IF NOT EXISTS rules (
    id SERIAL PRIMARY KEY,
    tractor_id INTEGER NOT NULL REFERENCES tractors(id),
    name TEXT NOT NULL DEFAULT '',
    field_name TEXT NOT NULL DEFAULT '',
    val_int INTEGER NOT NULL DEFAULT 0,
    val_float FLOAT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS notes(
    id SERIAL PRIMARY KEY,
    tractor_id INTEGER NOT NULL REFERENCES tractors(id),
    message TEXT NOT NULL DEFAULT '',
    time TEXT
);