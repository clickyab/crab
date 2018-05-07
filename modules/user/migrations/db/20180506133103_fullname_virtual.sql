
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD full_name VARCHAR(200)  GENERATED ALWAYS AS (CONCAT(first_name," ", last_name)) VIRTUAL COMMENT 'it is virtual field' AFTER last_name;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP COLUMN full_name;
