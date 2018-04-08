
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `creatives` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`campaign_id` INT(10) UNSIGNED NOT NULL,
`status` ENUM('pending', 'accepted', 'rejected') NOT NULL,
`type` ENUM('banner', 'native', 'vast') NOT NULL,
`url` VARCHAR(255) NOT NULL,
`attributes` TEXT NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT NULL,
`archived_at` DATETIME NULL DEFAULT NULL,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `assets` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`creative_id` INT(10) UNSIGNED NOT NULL,
`asset_type` ENUM('image', 'video', 'text', 'number') NOT NULL,
`property` TEXT NULL DEFAULT NULL,
`asset_key` VARCHAR(60) NULL DEFAULT NULL,
`asset_value` VARCHAR(255) NULL DEFAULT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE creatives
  ADD CONSTRAINT creative_campaign_id_fk
FOREIGN KEY (campaign_id) REFERENCES campaigns (id);

ALTER TABLE assets
  ADD CONSTRAINT assets_creative_id_fk
FOREIGN KEY (creative_id) REFERENCES creatives (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE assets;
DROP TABLE creatives;

