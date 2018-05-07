
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD advantage int unsigned DEFAULT 0 NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP advantage;

