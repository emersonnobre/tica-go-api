CREATE TABLE IF NOT EXISTS customers (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name VARCHAR(255) NULL,
  phone VARCHAR(20) NULL,
  cpf VARCHAR(11) NULL,
  email VARCHAR(255) NULL,
  instagram VARCHAR(255) NULL,
  birthday DATE NULL,
  created_at DATETIME NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
);