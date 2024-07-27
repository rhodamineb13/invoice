CREATE TABLE IF NOT EXISTS invoices(
    id BIGINT AUTO_INCREMENT NOT NULL,
    issue_date DATE NOT NULL,
    subject VARCHAR(255) NOT NULL,
    cust_id BIGINT NOT NULL,
    due_date DATE NOT NULL,
    status VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);