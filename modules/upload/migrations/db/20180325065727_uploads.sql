
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `uploads` (
`id` VARCHAR(191) NOT NULL,
`user_id` INT(10) UNSIGNED NOT NULL,
`mime` VARCHAR(50) NOT NULL,
`size` INT(11) NOT NULL,
`section` VARCHAR(50) NOT NULL,
`attr` VARCHAR(511) NULL DEFAULT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE uploads
  ADD CONSTRAINT uploads_users_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);

CREATE INDEX uploads_users_id_fk
  ON uploads (user_id);

ALTER TABLE users
  ADD CONSTRAINT users_uploads_id_fk
FOREIGN KEY (avatar) REFERENCES uploads (id);

ALTER TABLE domains
  ADD CONSTRAINT domains_uploads_id_fk
FOREIGN KEY (logo) REFERENCES uploads (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE uploads;

