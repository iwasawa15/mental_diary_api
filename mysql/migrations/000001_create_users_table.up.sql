CREATE TABLE IF NOT EXISTS users(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name CHAR (50) UNIQUE NOT NULL,
	email CHAR (100) UNIQUE NOT NULL
);