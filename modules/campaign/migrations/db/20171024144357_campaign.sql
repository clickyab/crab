
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaigns MODIFY end_at TIMESTAMP NULL DEFAULT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
