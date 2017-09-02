
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users MODIFY COLUMN status enum('registered', 'blocked','active') default 'registered' not null;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users MODIFY COLUMN status enum('registered', 'blocked') default 'registered' not null;
