-- Create the database
CREATE DATABASE IF NOT EXISTS lottery;

-- Use the database
USE lottery;

-- Table to store general users
CREATE TABLE IF NOT EXISTS general_users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
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
    number_id INT AUTO_INCREMENT PRIMARY KEY,
    number_value CHAR(6) UNIQUE NOT NULL
);

-- Table to manage available lottery numbers for each draw
CREATE TABLE IF NOT EXISTS draw_numbers (
    draw_id INT,
    number_id INT,
    available BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (draw_id, number_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store member purchases
CREATE TABLE IF NOT EXISTS purchases (
    purchase_id INT AUTO_INCREMENT PRIMARY KEY,
    member_id INT,
    number_id INT,
    purchase_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (member_id) REFERENCES members(member_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store transfers of lottery numbers
CREATE TABLE IF NOT EXISTS transfers (
    transfer_id INT AUTO_INCREMENT PRIMARY KEY,
    from_member_id INT,
    to_member_id INT,
    number_id INT,
    transfer_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    confirmed BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (from_member_id) REFERENCES members(member_id),
    FOREIGN KEY (to_member_id) REFERENCES members(member_id),
    FOREIGN KEY (number_id) REFERENCES lottery_numbers(number_id)
);

-- Table to store lottery draw results
CREATE TABLE IF NOT EXISTS lottery_results (
    draw_id INT AUTO_INCREMENT PRIMARY KEY,
    draw_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    number_1 CHAR(6),
    number_2 CHAR(6),
    number_3 CHAR(6),
    number_4 CHAR(6),
    number_5 CHAR(6),
    FOREIGN KEY (number_1) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_2) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_3) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_4) REFERENCES lottery_numbers(number_id),
    FOREIGN KEY (number_5) REFERENCES lottery_numbers(number_id)
);

-- Table to store system administrators
CREATE TABLE IF NOT EXISTS admins (
    admin_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Generate lottery numbers from 000000 to 999999
-- (This would normally be done with a script or program)
-- Example for MySQL, generating numbers using a stored procedure:

DELIMITER //

CREATE PROCEDURE generate_lottery_numbers()
BEGIN
    DECLARE num CHAR(6);
    SET num = '000000';
    
    WHILE num <= '999999' DO
        INSERT IGNORE INTO lottery_numbers (number_value) VALUES (num);
        SET num = LPAD(CAST(CAST(num AS UNSIGNED) + 1 AS CHAR), 6, '0');
    END WHILE;
END //

DELIMITER ;

-- Call the procedure to generate numbers
CALL generate_lottery_numbers();
