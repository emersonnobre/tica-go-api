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
