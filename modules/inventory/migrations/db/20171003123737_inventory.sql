
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `inventories` CHANGE `create_at` `created_at` TIMESTAMP NOT NULL;
ALTER TABLE `inventories` CHANGE `update_at` `updated_at` TIMESTAMP NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `inventories` CHANGE `created_at` `create_at` TIMESTAMP NOT NULL;
ALTER TABLE `inventories` CHANGE `updated_at` `update_at` TIMESTAMP NOT NULL;

