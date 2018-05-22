
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `domains` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`domain_base` VARCHAR(254) NOT NULL,
`attributes` TEXT NULL DEFAULT NULL,
`description` VARCHAR(254) NULL DEFAULT NULL,
`status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `users_domains` (
`domain_id` INT(10) UNSIGNED NOT NULL,
`user_id` INT(10) UNSIGNED NOT NULL,
`status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
`created_at` DATETIME NOT NULL,
`updated_at` DATETIME NULL,
PRIMARY KEY (`domain_id`, `user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE users_domains
  ADD CONSTRAINT domain_user_domains_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

ALTER TABLE users_domains
  ADD CONSTRAINT domain_user_users_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users_domains;
DROP TABLE domains;

