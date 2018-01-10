
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
SET foreign_key_checks = 0;;
DELETE FROM date_table WHERE id <= 10;
UPDATE date_table SET id=id-10;
DELETE FROM hour_table WHERE id <= 240;
UPDATE hour_table SET id=id-240,date_id=date_id-10;
SET foreign_key_checks = 1;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


