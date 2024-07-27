ALTER TABLE invoices
ADD CONSTRAINT FK_cust_id
FOREIGN KEY (cust_id) REFERENCES customers(id);