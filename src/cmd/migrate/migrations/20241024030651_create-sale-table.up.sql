CREATE TABLE IF NOT EXISTS type_of_payment (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  description VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS sale (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  total_price FLOAT NOT NULL,
  discount FLOAT NULL,
  comments VARCHAR(255) NULL,
  type_of_payment_id INT NOT NULL,
  created_at DATETIME NOT NULL,
  employee_id INT NOT NULL,
  customer_id INT NOT NULL,
  FOREIGN KEY (type_of_payment_id) REFERENCES type_of_payment(id),
  FOREIGN KEY (employee_id) REFERENCES employees(id),
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS sale_item (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  quantity INT NOT NULL,
  product_id INT NOT NULL,
  sale_id INT NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (sale_id) REFERENCES sale(id)
);