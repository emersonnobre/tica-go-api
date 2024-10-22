ALTER TABLE transactions
ADD COLUMN created_by INT NOT NULL,
ADD CONSTRAINT fk_created_by
FOREIGN KEY (created_by)
REFERENCES employee(id);