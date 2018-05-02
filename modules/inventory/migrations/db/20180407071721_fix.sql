
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE publishers CHANGE COLUMN `publisher` `supplier` VARCHAR(127) NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE publishers CHANGE COLUMN `supplier` `publisher` VARCHAR(127) NOT NULL;


