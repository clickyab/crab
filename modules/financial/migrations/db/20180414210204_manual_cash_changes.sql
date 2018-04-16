
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `manual_cash_changes`;

CREATE TABLE `manual_cash_changes` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `operator_id` INT(10) UNSIGNED NOT NULL,
  `reason` ENUM('gift', 'manual_pay', 'refund') NOT NULL,
  `amount` INT(10) NOT NULL,
  `description` TEXT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `manual_cash_changes`;

