--通过 SHOW VARIABLES LIKE 'char%' 查看支持的编码方式 
CREATE DATABASE IF NOT EXISTS douro DEFAULT CHARACTER SET 'utf8';

USE douro;

--创建用户信息
CREATE TABLE IF NOT EXISTS user_profile(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    gender TINYINT NOT NULL DEFAULT 0 COMMENT '0:保密,1:男,2:女' ,
    phone VARCHAR(11) NOT NULL,
    avatar VARCHAR(100) NOT NULL DEFAULT '/images/avatar/default.png',
    register_mode TINYINT NOT NULL DEFAULT 0 COMMENT '0:电话,1:微信',
    third_party_id VARCHAR(50) NULL,
    create_time DATETIME  DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME  DEFAULT CURRENT_TIMESTAMP
);

--创建账户表
CREATE TABLE IF NOT EXISTS user_auth(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    encrypt_pwd VARCHAR(128) NOT NULL,
    user_id INT UNSIGNED NOT NULL UNIQUE,
    create_time DATETIME  DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME  DEFAULT CURRENT_TIMESTAMP
);



