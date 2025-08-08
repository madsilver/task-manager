-- +migrate Down
DELETE FROM users WHERE email IN ('manager@company.com', 'tech1@company.com');
