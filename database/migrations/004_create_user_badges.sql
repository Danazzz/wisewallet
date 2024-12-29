CREATE TABLE user_badges (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    badge_id INT REFERENCES badges(id)
);