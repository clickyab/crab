
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `roles` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`name` VARCHAR(40) NOT NULL,
`description` VARCHAR(255) NULL DEFAULT NULL,
`domain_id` INT(10) UNSIGNED NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `role_permission` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`role_id` INT(10) UNSIGNED NOT NULL,
`perm` VARCHAR(60) NOT NULL,
`scope` ENUM('self', 'global') NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
INDEX `role_permission_roles_idx` (`role_id` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `role_user` (
`user_id` INT(10) UNSIGNED NOT NULL,
`role_id` INT(10) UNSIGNED NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`, `role_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE advisors
  ADD CONSTRAINT parent2_user_users_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE advisors
  ADD CONSTRAINT advisors_users_id_fk
FOREIGN KEY (advisor_id) REFERENCES users (id);

ALTER TABLE role_user
  ADD CONSTRAINT role_user_users_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE role_permission
  ADD CONSTRAINT role_permission_roles_id_fk
FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE role_user
  ADD CONSTRAINT role_user_roles_id_fk
FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE advisors
  ADD CONSTRAINT advisors_domains_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

ALTER TABLE roles
  ADD CONSTRAINT roles_domains_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE role_user;
DROP TABLE role_permission;
DROP TABLE roles;

