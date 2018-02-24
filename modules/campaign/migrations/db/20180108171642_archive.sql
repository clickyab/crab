
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaigns ADD archive_at TIMESTAMP DEFAULT NULL  NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaigns DROP archive_at;

