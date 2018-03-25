
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `inventories` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` INT(10) UNSIGNED NOT NULL,
`label` VARCHAR(60) NOT NULL COMMENT 'user personal label',
`status` ENUM('disable', 'enable') NOT NULL,
`publisher_type` ENUM('web', 'app') NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `publishers` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(127) NOT NULL,
  `domain` VARCHAR(127) NOT NULL,
  `categories` TEXT NULL DEFAULT NULL,
  `publisher` VARCHAR(127) NOT NULL,
  `kind` ENUM('app', 'web') NOT NULL,
  `status` ENUM('pending', 'blocked', 'accepted') NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `domain` (`domain` ASC))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `inventories_publishers` (
  `publisher_id` INT(10) UNSIGNED NOT NULL,
  `inventory_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`publisher_id`, `inventory_id`))
  ENGINE = InnoDB;

CREATE INDEX inventory_status_update_user_id
  ON inventories (user_id, status, updated_at);

ALTER TABLE inventories_publishers
  ADD CONSTRAINT fk_publishers
FOREIGN KEY (publisher_id) REFERENCES publishers (id);

ALTER TABLE inventories_publishers
  ADD CONSTRAINT fk_inventories
FOREIGN KEY (inventory_id) REFERENCES inventories (id);


ALTER TABLE inventories
  ADD CONSTRAINT inventory_user_fk
FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE campaigns
  ADD CONSTRAINT `campaign_inventory_id_fk` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`);



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE inventories_publishers;
DROP TABLE publishers;
DROP TABLE inventories;

