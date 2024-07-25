CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    nickname VARCHAR UNIQUE,
    password VARCHAR -- TODO: change to hash
);

CREATE TABLE short_urls(
    id VARCHAR PRIMARY KEY, -- TODO: special UUID type
    original_url VARCHAR,
    user_id INT NOT NULL REFERENCES users(id)
);
