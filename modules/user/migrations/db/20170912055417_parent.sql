
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE consular_customer;
CREATE TABLE `parent_user` (
  `user_id` int(11) NOT NULL,
  `parent_id` int(11) NOT NULL,
  `domain_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`parent_id`,`domain_id`),
  KEY `parent_user_users_id_fk` (`parent_id`),
  KEY `parent_user_domains_id_fk` (`domain_id`),
  CONSTRAINT `parent2_user_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `parent_user_domains_id_fk` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`),
  CONSTRAINT `parent_user_users_id_fk` FOREIGN KEY (`parent_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE parent_user;
CREATE TABLE consular_customer
(
  consulary_id INT NOT NULL,
  customer_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  PRIMARY KEY (consulary_id,customer_id),
  CONSTRAINT cons_fk FOREIGN KEY (consulary_id) REFERENCES users (id),
  CONSTRAINT cost_fk FOREIGN KEY (customer_id) REFERENCES users (id)
);

