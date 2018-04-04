
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `campaigns` ADD `tld` VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'top level domain' AFTER `title`;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `campaigns` DROP `tld`;