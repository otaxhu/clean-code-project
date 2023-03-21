DROP DATABASE IF EXISTS clean_code;
CREATE DATABASE clean_code;
USE clean_code;
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);