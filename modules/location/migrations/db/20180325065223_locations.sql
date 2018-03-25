
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `countries` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`name` VARCHAR(50) NOT NULL,
PRIMARY KEY (`id`),
UNIQUE INDEX `countries_name_uidx` (`name` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `provinces` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `country_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `provinces_name_country_id_unidex` (`name` ASC, `country_id` ASC))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `cities` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `province_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `cities_name_provinces_id_unidex` (`name` ASC, `province_id` ASC))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

ALTER TABLE cities
  ADD CONSTRAINT cities_provinces_id_fk
FOREIGN KEY (province_id) REFERENCES provinces (id);

ALTER TABLE provinces
  ADD CONSTRAINT provinces_countries_id_fk FOREIGN KEY (country_id) REFERENCES countries (id);

ALTER TABLE users
  ADD CONSTRAINT user_city_id_fk
FOREIGN KEY (city_id) REFERENCES cities (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cities;
DROP TABLE provinces;
DROP TABLE countries;

