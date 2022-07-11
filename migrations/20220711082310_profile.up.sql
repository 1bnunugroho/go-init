CREATE TABLE profile
(
    id         VARCHAR PRIMARY KEY,
    bio        VARCHAR NULL,
    image      VARCHAR NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);