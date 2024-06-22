CREATE DATABASE IF NOT EXISTS test;

USE test;

CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    UNIQUE KEY unique_name_email (name, email)
);

INSERT IGNORE INTO user (name, email) VALUES ('TestUser1', 'test1@example.com');
INSERT IGNORE INTO user (name, email) VALUES ('TestUser2', 'test2@example.com');
INSERT IGNORE INTO user (name, email) VALUES ('TestUser3', 'test3@example.com');