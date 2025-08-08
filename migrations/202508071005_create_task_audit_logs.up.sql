-- +migrate Up
CREATE TABLE task_audit_logs (
     id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
     task_id BIGINT NOT NULL,
     operation ENUM('CREATE', 'UPDATE', 'DELETE', 'READ') NOT NULL,
     performed_by BIGINT NULL,
     performed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
