CREATE TABLE uzer
(
    id         VARCHAR PRIMARY KEY,
    email       VARCHAR NOT NULL,
    username   VARCHAR NOT NULL,
    password   VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);