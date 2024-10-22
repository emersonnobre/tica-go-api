CREATE TABLE IF NOT EXISTS transactions (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  reason VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  `type` INT NOT NULL,
  created_at DATETIME NOT NULL,
  product_id INT NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id)
);