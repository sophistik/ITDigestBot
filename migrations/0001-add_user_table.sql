-- +migrate Up
CREATE TABLE users 
(
    id SERIAL NOT NULL PRIMARY KEY,
    tags TEXT[],
    last_input TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +migrate Down
DROP TABLE users;