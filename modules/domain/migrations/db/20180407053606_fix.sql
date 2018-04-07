
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users_domains MODIFY COLUMN created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE users_domains MODIFY COLUMN updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


