CREATE TABLE IF NOT EXISTS sale_item (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  quantity INT NOT NULL,
  product_id INT NOT NULL,
  sale_id INT NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (sale_id) REFERENCES sale(id)
);