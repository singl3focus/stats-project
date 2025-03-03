CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE stats (
    user_id INT REFERENCES users(id),
    service_id INT REFERENCES services(id),
    count INT NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, service_id)
);