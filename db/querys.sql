CREATE DATABASE IF NOT EXISTS vendasdb;
use vendasdb;
---- TABELAS ---
	-- ID        uint64   
	-- Name      string    
	-- Cpf       string    
	-- Type      uint      
	-- Phone     string    
	-- Password  string    
	-- Email     string    
	-- CreatedAt time.Time 
	-- UpdatedAt time.Time 

DROP TABLE IF EXISTS users;

--Cria a tabela users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name varchar(50) not null,
    cpf varchar(11) not null unique,
    type int not null,
    phone varchar(11) not null unique,
    password varchar(100) not null,
    email varchar(50) not null unique,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp
);