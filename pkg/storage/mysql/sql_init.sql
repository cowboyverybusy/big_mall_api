-- 创建cowboy_mall数据库
CREATE DATABASE IF NOT EXISTS cowboy_mall
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- 使用cowboy_mall数据库
USE cowboy_mall;

-- 创建users表
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INT,
    status TINYINT NOT NULL DEFAULT 1 COMMENT '用户状态: 1-活跃, 0-停用',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX idx_username (username),
    UNIQUE INDEX idx_email (email),
    INDEX idx_deleted_at (deleted_at)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;