CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    gender VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    password VARCHAR
);