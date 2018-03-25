
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`email` VARCHAR(50) NOT NULL,
`password` VARCHAR(60) NOT NULL,
`access_token` VARCHAR(60) NOT NULL,
`avatar` VARCHAR(191) NULL DEFAULT NULL,
`status` ENUM('registered', 'blocked', 'active') NOT NULL DEFAULT 'registered',
`old_password` VARCHAR(255) NULL DEFAULT NULL,
`city_id` INT(10) UNSIGNED NULL DEFAULT NULL,
`ssn` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT 'Social Security Number',
`land_line` VARCHAR(20) NULL DEFAULT NULL,
`cellphone` VARCHAR(20) NULL DEFAULT NULL,
`postal_code` VARCHAR(10) NULL DEFAULT NULL,
`first_name` VARCHAR(40) NOT NULL,
`last_name` VARCHAR(40) NOT NULL,
`address` VARCHAR(255) NULL DEFAULT NULL,
`gender` ENUM('male', 'female', 'not_specified') NOT NULL,
`attributes` TEXT NULL DEFAULT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE INDEX `users_email_uindex` (`email` ASC),
UNIQUE INDEX `users_token_uindex` (`access_token` ASC),
UNIQUE INDEX `ssn` (`ssn` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


CREATE TABLE IF NOT EXISTS `corporations` (
  `user_id` INT(10) UNSIGNED NOT NULL,
  `legal_name` VARCHAR(50) NOT NULL,
  `legal_register` VARCHAR(50) NULL DEFAULT NULL,
  `economic_code` VARCHAR(50) NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `advisors` (
  `user_id` INT(10) UNSIGNED NOT NULL,
  `advisor_id` INT(10) UNSIGNED NOT NULL,
  `domain_id` INT(10) UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `advisor_id`, `domain_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

ALTER TABLE corporations
  ADD CONSTRAINT corporations_user_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);

CREATE INDEX user_city_id_fk
  ON users (city_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE corporations;
DROP TABLE advisors;
DROP TABLE users;

