
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users MODIFY advantage int;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users MODIFY advantage int not null default 0;

