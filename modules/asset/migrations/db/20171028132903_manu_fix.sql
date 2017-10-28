
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE manufacturers CHANGE `brand` `name` VARCHAR(191) NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE manufacturers CHANGE `name` `brand` VARCHAR(191) NOT NULL;


