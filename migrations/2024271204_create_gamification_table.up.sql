CREATE TABLE gamification (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    points INT DEFAULT 0,
    badges TEXT[],
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);