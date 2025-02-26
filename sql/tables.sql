DROP DATABASE IF EXISTS devbook;
CREATE DATABASE devbook;

Drop table if exists usuarios;

CREATE TABLE usuarios (
	id INT PRIMARY KEY,
	nome VARCHAR(100) NOT NULL,
	nick VARCHAR(50) NOT NULL unique,
	email VARCHAR(100) NOT NULL,
	senha VARCHAR(255) NOT NULL,
	criado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);