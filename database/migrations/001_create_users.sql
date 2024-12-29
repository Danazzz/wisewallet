CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Insert default admin user
-- Password 'admin' hashed memakai bcrypt
INSERT INTO users (username, password)
VALUES ('admin', '$2a$10$CwTycUXWue0Thq9StjUM0uJ8QlAv3Myc2m1MvBOQ.k6TrHDpQG2yW');