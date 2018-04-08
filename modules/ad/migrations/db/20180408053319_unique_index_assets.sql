
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `assets` ADD UNIQUE `unique_asset` (`creative_id`, `asset_key`);



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE `assets` DROP INDEX unique_asset;
