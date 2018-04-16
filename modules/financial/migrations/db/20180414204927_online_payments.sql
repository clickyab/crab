
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `online_payments`;
CREATE TABLE `online_payments` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `gateway_id` INT(10) UNSIGNED NOT NULL,
  `amount` INT(10) UNSIGNED NOT NULL,
  `status` ENUM('init', 'in_gateway', 'back_to_clickyab', 'finalized') NOT NULL,
  `bank_status` INT(10) UNSIGNED NOT NULL,
  `ref_num` VARCHAR(45) NOT NULL,
  `res_num` VARCHAR(45) NULL,
  `cid` VARCHAR(45) NULL,
  `trace_number` VARCHAR(45) NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `ref_num_UNIQUE` (`ref_num` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `online_payments`;

