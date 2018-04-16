
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `bank_snaps`;

CREATE TABLE `bank_snaps` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `trace_number` INT(10) UNSIGNED NOT NULL,
  `vat` INT(10) NOT NULL,
  `amount` INT(10) NOT NULL,
  `pay_amount` INT(10) NOT NULL,
  `status` ENUM('pending', 'accepted', 'rejected') NOT NULL,
  `checked_by` INT(10) UNSIGNED NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `bank_snaps`;

