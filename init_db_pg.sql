-- Drop existing tables if they exist
-- DROP TABLE IF EXISTS lottery_results CASCADE;
-- DROP TABLE IF EXISTS transfers CASCADE;
-- DROP TABLE IF EXISTS purchases CASCADE;
-- DROP TABLE IF EXISTS draw_numbers CASCADE;
-- DROP TABLE IF EXISTS members CASCADE;
-- DROP TABLE IF EXISTS general_users CASCADE;
-- DROP TABLE IF EXISTS admins CASCADE;
-- DROP SEQUENCE IF EXISTS lottery_number_seq;

-- Create tables
-- Table to store general users
CREATE TABLE IF NOT EXISTS general_users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_member BOOLEAN DEFAULT FALSE
);

-- Table to store members (inherits general user properties)
CREATE TABLE IF NOT EXISTS members (
    member_id INT PRIMARY KEY,
    wallet_balance DECIMAL(10, 2) DEFAULT 0.00,
    FOREIGN KEY (member_id) REFERENCES general_users(user_id)
);

-- Table to store lottery numbers
CREATE TABLE IF NOT EXISTS lottery_numbers (
    number_id VARCHAR(10) PRIMARY KEY,
    number_value CHAR(6) UNIQUE NOT NULL
);

-- Table to manage available lottery numbers for each draw
CREATE TABLE IF NOT EXISTS draw_numbers (
    draw_id INT,
    number_id VARCHAR(10),
    available BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (draw_id, number_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store member purchases
CREATE TABLE IF NOT EXISTS purchases (
    purchase_id SERIAL PRIMARY KEY,
    member_id INT,
    number_id VARCHAR(10),
    purchase_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (member_id) REFERENCES members(member_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store transfers of lottery numbers
CREATE TABLE IF NOT EXISTS transfers (
    transfer_id SERIAL PRIMARY KEY,
    from_member_id INT,
    to_member_id INT,
    number_id VARCHAR(10),
    transfer_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    confirmed BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (from_member_id) REFERENCES members(member_id),
    FOREIGN KEY (to_member_id) REFERENCES members(member_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store lottery draw results
CREATE TABLE IF NOT EXISTS lottery_results (
    draw_id SERIAL PRIMARY KEY,
    draw_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    number_1 VARCHAR(10),
    number_2 VARCHAR(10),
    number_3 VARCHAR(10),
    number_4 VARCHAR(10),
    number_5 VARCHAR(10),
    FOREIGN KEY (number_1) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_2) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_3) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_4) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_5) REFERENCES lottery_numbers(number_id)
);

-- Table to store system administrators
CREATE TABLE IF NOT EXISTS admins (
    admin_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Create a sequence for generating unique lottery numbers
CREATE SEQUENCE IF NOT EXISTS lottery_number_seq
START 1
INCREMENT 1
MINVALUE 1
MAXVALUE 999999
CACHE 1;

-- Function to generate a specified number of lottery numbers
CREATE OR REPLACE FUNCTION generate_lottery_numbers(num_count INTEGER)
RETURNS void LANGUAGE plpgsql AS $$
DECLARE
    i INTEGER := 0;
    next_seq_value INT;
    formatted_num CHAR(6);
    random_num INT;
    lottery_id TEXT;
    lottery_number TEXT;
    affected_rows INTEGER; -- Changed to INTEGER to capture the row count
BEGIN
    WHILE i < num_count LOOP
    -- Generate a random 6-digit number
        random_num := FLOOR(RANDOM() * 1000000)::INT;
        lottery_number := LPAD(random_num::TEXT, 6, '0');
        
        -- Get the next value from the sequence
        next_seq_value := NEXTVAL('lottery_number_seq');
        formatted_num := LPAD(next_seq_value::TEXT, 6, '0');
        lottery_id := CONCAT('LT-', formatted_num);

        -- Insert into the table if the lottery_id does not already exist
        INSERT INTO lottery_numbers (number_id, number_value)
        VALUES (lottery_id, lottery_number)
        ON CONFLICT (number_id) DO NOTHING;

        -- Check if the insert was successful
        GET DIAGNOSTICS affected_rows = ROW_COUNT;

        -- If a row was inserted, increment counter
        IF affected_rows > 0 THEN
            i := i + 1;
        END IF;
    END LOOP;
END;
$$;

-- Example: Generate 100 lottery numbers
-- Uncomment the following line to populate the lottery numbers
SELECT generate_lottery_numbers(100);
