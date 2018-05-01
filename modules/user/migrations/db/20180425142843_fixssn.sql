
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE users SET ssn=NULL;
ALTER TABLE users MODIFY ssn CHAR(10) COMMENT 'Social Security Number';

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
UPDATE users SET ssn=NULL;
ALTER TABLE users MODIFY ssn INT UNSIGNED COMMENT 'Social Security Number';

