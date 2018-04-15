
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `billings`;
CREATE TABLE `billings` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `pay_model` ENUM('online_payment', 'bank_snap', 'manual_cash_change') NOT NULL,
  `income_id` INT(10) UNSIGNED NOT NULL,
  `vat` INT(10) NOT NULL,
  `amount` INT(10) NOT NULL,
  `pay_amount` INT(10) NOT NULL,
  `balance` INT(10) NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `billings`;

