-- +migrate Up
INSERT INTO users (name, email, password, roles) VALUES
('manager', 'manager@company.com', '$2a$10$JgL/u7.zEERMOz2pES7bAePpXJYmjlUAx3iXOj0TNcnSl7183XExK', JSON_ARRAY('manager')),
('tech1', 'tech1@company.com', '$2a$10$JgL/u7.zEERMOz2pES7bAePpXJYmjlUAx3iXOj0TNcnSl7183XExK', JSON_ARRAY('technician'));
