
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `creatives` ADD `max_bid` INT(10) UNSIGNED NOT NULL AFTER `url`;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `creatives` DROP `max_bid`;
