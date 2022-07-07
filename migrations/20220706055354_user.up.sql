CREATE TABLE uzer
(
    id         VARCHAR PRIMARY KEY,
    email       VARCHAR NOT NULL,
    username   VARCHAR NOT NULL,
    password   VARCHAR NOT NULL,
    bio        VARCHAR NULL,
    image      VARCHAR NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);