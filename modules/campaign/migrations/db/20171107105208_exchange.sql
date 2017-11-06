
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaigns ADD exchange BOOLEAN NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaigns DROP COLUMN exchange BOOLEAN NOT NULL;


