-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

DELETE FROM user_personal;
DELETE FROM user_corporation;
DELETE FROM domain_user;
DELETE FROM consular_customer;
DELETE FROM user_wlbl_presets;
DELETE FROM uploads;
DELETE FROM role_user;
DELETE FROM users;

ALTER TABLE users DROP user_type;

ALTER TABLE users ADD city_id INT NULL;
ALTER TABLE users ADD CONSTRAINT user_city_id_fk FOREIGN KEY (city_id) REFERENCES cities (id);

ALTER TABLE users ADD ssn INT NULL COMMENT 'Social Security Number';
ALTER TABLE users ADD land_line VARCHAR(20) NULL;
ALTER TABLE users ADD cellphone VARCHAR(20) NULL;
ALTER TABLE users ADD postal_code VARCHAR(10) NULL;
ALTER TABLE users ADD first_name VARCHAR(40) NOT NULL;
ALTER TABLE users ADD last_name VARCHAR(40) NOT NULL;
ALTER TABLE users ADD address VARCHAR(255) NULL;
ALTER TABLE users ADD gender ENUM ('male', 'female', 'not_specified') NOT NULL;

CREATE TABLE corporations (
  id             INT AUTO_INCREMENT PRIMARY KEY,
  user_id        INT         NOT NULL,
  legal_name     VARCHAR(50) NOT NULL,
  legal_register VARCHAR(50) NULL,
  economic_code  VARCHAR(50) NULL,
  CONSTRAINT corporations_user_id_uindex
  UNIQUE (user_id),
  CONSTRAINT corporations_user_id_fk
  FOREIGN KEY (user_id) REFERENCES users (id)
);

DROP TABLE user_personal;
DROP TABLE user_corporation;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM corporations;
DELETE FROM domain_user;
DELETE FROM consular_customer;
DELETE FROM user_wlbl_presets;
DELETE FROM uploads;
DELETE FROM role_user;
DELETE FROM users;

CREATE TABLE `user_personal` (
  `user_id`     INT(11)     NOT NULL,
  `first_name`  VARCHAR(40) NOT NULL,
  `last_name`   VARCHAR(40) NOT NULL,
  `gender`      ENUM ('male', 'female', 'not_specified') DEFAULT NULL,
  `cellphone`   VARCHAR(20)                              DEFAULT NULL,
  `phone`       VARCHAR(20)                              DEFAULT NULL,
  `address`     VARCHAR(255)                             DEFAULT NULL,
  `city_id`     INT(11)                                  DEFAULT NULL,
  `created_at`  TIMESTAMP   NOT NULL                     DEFAULT NOT NULL,
  `updated_at`  TIMESTAMP   NOT NULL                     DEFAULT NOT NULL,
  `province_id` INT(11)                                  DEFAULT NULL,
  `country_id`  INT(11)                                  DEFAULT NULL,
  `zip_code`    INT(11)                                  DEFAULT NULL,
  `national_id` VARCHAR(30)                              DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_personal_user_id_uindex` (`user_id`),
  CONSTRAINT `user_personal_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `user_corporation` (
  `user_id`       INT(11)     NOT NULL,
  `name`          VARCHAR(50) NOT NULL,
  `cellphone`     VARCHAR(20)  NULL,
  `phone`         VARCHAR(20)  NULL,
  `address`       VARCHAR(255) NULL,
  `economic_code` VARCHAR(40)  NULL,
  `register_code` VARCHAR(40)  NULL,
  `city_id`       INT(11)      NULL,
  `created_at`    TIMESTAMP    NOT NULL,
  `updated_at`    TIMESTAMP    NOT NULL,
  `last_name`     VARCHAR(40)  NOT NULL,
  `first_name`    VARCHAR(40)  NOT NULL,
  `province_id`   INT(11)      DEFAULT NULL,
  `country_id`    INT(11)      DEFAULT NULL,
  `zip_code`      INT(11)      DEFAULT NULL,
  `national_id`   VARCHAR(30)  DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_corporation_user_id_uindex` (`user_id`),
  CONSTRAINT `user_corporation_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

ALTER TABLE users ADD user_type enum('personal','corporation') NOT NULL;

ALTER TABLE users DROP city_id;
ALTER TABLE users DROP FOREIGN KEY user_city_id_fk;

ALTER TABLE users DROP ssn;
ALTER TABLE users DROP land_line;
ALTER TABLE users DROP cellphone;
ALTER TABLE users DROP postal_code;
ALTER TABLE users DROP first_name;
ALTER TABLE users DROP last_name;
ALTER TABLE users DROP address;
ALTER TABLE users DROP gender;


DROP TABLE corporations;
