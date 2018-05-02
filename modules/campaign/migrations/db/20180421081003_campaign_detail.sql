
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `campaign_detail` ADD COLUMN IF NOT EXISTS `publisher_id` INT UNSIGNED NOT NULL ;
ALTER TABLE `campaign_detail` ADD CONSTRAINT `campaign_detail_publisher_id_fk` FOREIGN KEY (`publisher_id`)
  REFERENCES `publishers` (`id`);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaign_detail DROP FOREIGN KEY `campaign_detail_publisher_id_fk`;

