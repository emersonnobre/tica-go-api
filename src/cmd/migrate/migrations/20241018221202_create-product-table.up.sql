CREATE TABLE IF NOT EXISTS products (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  purchase_price FLOAT NOT NULL,
  sale_price FLOAT NOT NULL,
  stock INT NOT NULL,
  barcode VARCHAR(255) NOT NULL,
  category_id INT NOT NULL,
  active BOOLEAN NOT NULL,
  created_at DATETIME NOT NULL,
  created_by INT NOT NULL,
  updated_at DATETIME NULL,
  updated_by INT NULL,
  is_feedstock BOOLEAN NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id),
  FOREIGN KEY (created_by) REFERENCES employees(id),
  FOREIGN KEY (updated_by) REFERENCES employees(id)
);