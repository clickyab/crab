
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
SET FOREIGN_KEY_CHECKS=0;
ALTER TABLE `roles` DROP FOREIGN KEY IF EXISTS `roles_domains_id_fk`;
ALTER TABLE roles DROP IF EXISTS domain_id;
ALTER TABLE roles ADD level int DEFAULT 1 NOT NULL;
TRUNCATE TABLE roles;
SET FOREIGN_KEY_CHECKS=1;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


