ALTER TABLE transactions
ADD COLUMN created_by INT NOT NULL,
ADD FOREIGN KEY fk_created_by(created_by) REFERENCES employees(id);