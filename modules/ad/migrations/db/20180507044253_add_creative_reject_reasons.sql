
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `creatives` ADD `reject_reasons_id` INT(10) UNSIGNED NULL AFTER `attributes`;
CREATE TABLE  IF NOT EXISTS `creative_reject_reasons` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`reason` TEXT NOT NULL,
`status` ENUM('enable', 'disable') NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT NULL,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE creatives
  ADD CONSTRAINT reject_reasons_id_fk
FOREIGN KEY (reject_reasons_id) REFERENCES creative_reject_reasons (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `creatives` DROP FOREIGN KEY `reject_reasons_id_fk`;
ALTER TABLE `creatives` DROP `reject_reasons_id`;
DROP TABLE `creative_reject_reasons`;