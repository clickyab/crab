
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DELETE FROM uploads;
ALTER TABLE uploads DROP COLUMN `path`;
ALTER TABLE uploads ADD COLUMN `attr` VARCHAR(511);
ALTER TABLE uploads MODIFY id VARCHAR(191) NOT NULL;
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE uploads DROP COLUMN `attr`;
ALTER TABLE uploads MODIFY id INT(11) NOT NULL;
ALTER TABLE uploads ADD COLUMN `path` VARCHAR(255);


