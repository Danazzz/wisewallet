CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Insert default admin user
-- Password 'admin' hashed memakai bcrypt
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin') THEN
        INSERT INTO users (username, password) 
        VALUES ('admin', '$2a$10$i4aBpDxq76xeaK8DOiAs4ugsvctO22RpUfwexuuBKQ36US1jhZXYe');
    END IF;
END $$;