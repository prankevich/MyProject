CREATE TABLE IF NOT EXISTS employees
(
    id SERIAL PRIMARY KEY,
    age   int NOT NULL,
    name  VARCHAR NOT NULL ,
    email VARCHAR NOT NULL unique
);