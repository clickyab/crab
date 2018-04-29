
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `audit_logs`;

CREATE TABLE `audit_logs` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) UNSIGNED DEFAULT NULL,
  `user_id` int(11) UNSIGNED DEFAULT NULL,
  `user_perm` VARCHAR(60) NOT NULL,
  `perm_scope` ENUM('self', 'global') NOT NULL,
  `action` ENUM('insert', 'update', 'delete') DEFAULT NULL,
  `target_model` varchar(255) NOT NULL,
  `target_id` int(11) UNSIGNED DEFAULT NULL,
  `owner_id` int(11) UNSIGNED DEFAULT NULL,
  `impersonate` TINYINT(1) UNSIGNED DEFAULT 0,
  `impersonator_id` int(11) UNSIGNED DEFAULT NULL,
  `description` text,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

DROP TABLE IF EXISTS `audit_log_details`;

CREATE TABLE `audit_log_details` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `audit_log_id` int(11) UNSIGNED NOT NULL,
  `data` text,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

ALTER TABLE audit_logs
  ADD CONSTRAINT audit_logs_domain_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

ALTER TABLE audit_logs
  ADD CONSTRAINT audit_logs_user_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE audit_logs
  ADD CONSTRAINT audit_logs_owner_id_fk
FOREIGN KEY (owner_id) REFERENCES users (id);

ALTER TABLE audit_logs
  ADD CONSTRAINT audit_logs_impersonator_id_fk
FOREIGN KEY (impersonator_id) REFERENCES users (id);

ALTER TABLE audit_log_details
  ADD CONSTRAINT audit_log_details_audit_id_fk
FOREIGN KEY (audit_log_id) REFERENCES audit_logs (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE audit_logs DROP FOREIGN KEY audit_logs_domain_id_fk;
ALTER TABLE audit_logs DROP FOREIGN KEY audit_logs_user_id_fk;
ALTER TABLE audit_logs DROP FOREIGN KEY audit_logs_owner_id_fk;
ALTER TABLE audit_logs DROP FOREIGN KEY audit_logs_impersonator_id_fk;
ALTER TABLE audit_log_details DROP FOREIGN KEY audit_log_details_audit_id_fk;

DROP TABLE `audit_logs`;
DROP TABLE `audit_log_details`;