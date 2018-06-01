
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE users_domains;
CREATE TABLE `users_domains` (
  `id` INT(10) AUTO_INCREMENT,
  `domain_id` int(10) unsigned,
  `role_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `status` enum('disable','enable') NOT NULL DEFAULT 'enable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `domain_user_users_id_fk` (`user_id`),
  CONSTRAINT `domain_user_domains_id_fk` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`),
  CONSTRAINT `domain_user_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `domain_user_roles_id_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


