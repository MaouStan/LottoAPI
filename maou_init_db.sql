-- Drop Tables if They Exist (for Resetting the Database)
DROP TABLE IF EXISTS prizes CASCADE;
DROP TABLE IF EXISTS draws CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS lotto_numbers CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Drop ENUM Types if They Exist (for Resetting the Database)
DO $$ BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        DROP TYPE user_role;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
        DROP TYPE transaction_type;
    END IF;
END $$;

-- Create ENUM Type for User Roles
CREATE TYPE user_role AS ENUM ('member', 'admin');

-- Create ENUM Type for Transaction Types
CREATE TYPE transaction_type AS ENUM ('purchase', 'claim_winnings');

-- Create Users Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role user_role DEFAULT 'member',
    wallet_balance INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Lotto Numbers Table
CREATE TABLE IF NOT EXISTS lotto_numbers (
    id SERIAL PRIMARY KEY,
    number VARCHAR(6) UNIQUE NOT NULL,
    is_sold BOOLEAN DEFAULT FALSE,
    sold_to INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Transactions Table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    lotto_number_id INT REFERENCES lotto_numbers(id) ON DELETE CASCADE,
    transaction_type transaction_type NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Draws Table
CREATE TABLE IF NOT EXISTS draws (
    id SERIAL PRIMARY KEY,
    draw_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    winning_numbers VARCHAR(6)[]
);

-- Create Prizes Table
CREATE TABLE IF NOT EXISTS prizes (
    id SERIAL PRIMARY KEY,
    draw_id INT REFERENCES draws(id) ON DELETE CASCADE,
    lotto_number_id INT REFERENCES lotto_numbers(id),
    prize_amount INTEGER NOT NULL,
    claimed BOOLEAN DEFAULT FALSE
);

-- Insert Admin User
INSERT INTO users (username, password_hash, role, wallet_balance)
VALUES ('goblin123', '$2b$10$IKO4jTzz2NxYFHE3bVfBveBuv7wDTUsY.57C2jt5VE02WIN99YT9a', 'admin', 0);

-- Insert 100 Random Lotto Numbers
DO $$
DECLARE
    i INT;
    random_number VARCHAR(6);
BEGIN
    FOR i IN 1..100 LOOP
        random_number := LPAD(CAST(FLOOR(RANDOM() * 1000000) AS VARCHAR), 6, '0');
        INSERT INTO lotto_numbers (number) VALUES (random_number);
    END LOOP;
END $$;
