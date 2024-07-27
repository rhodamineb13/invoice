ALTER TABLE orders
ADD CONSTRAINT FK_item_id FOREIGN KEY (item_id) REFERENCES items(id),
ADD CONSTRAINT FK_invoice_id FOREIGN KEY (invoice_id) REFERENCES invoices(id);