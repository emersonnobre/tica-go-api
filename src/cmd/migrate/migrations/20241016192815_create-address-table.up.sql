CREATE TABLE IF NOT EXISTS addresses (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  street VARCHAR(255) NOT NULL,
  neighborhood VARCHAR(255) NOT NULL,
  cep VARCHAR(255) NULL,
  customer_id INT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);