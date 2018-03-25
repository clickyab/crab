
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `browsers` (
`name` VARCHAR(60) NOT NULL,
`deleted_at` DATETIME NULL DEFAULT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`name`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `categories` (
  `name` VARCHAR(15) NOT NULL,
  `description` VARCHAR(300) NOT NULL,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`name`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `isps` (
  `name` VARCHAR(40) NOT NULL,
  `kind` ENUM('cellular', 'isp', 'both') NOT NULL,
  `status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `manufacturers` (
  `name` VARCHAR(191) NOT NULL,
  `status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`),
  UNIQUE INDEX `manufacturer_brand_uindex` (`name` ASC))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `oses` (
  `name` VARCHAR(40) NOT NULL,
  `status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `platforms` (
  `name` VARCHAR(15) NOT NULL,
  `status` ENUM('disable', 'enable') NOT NULL DEFAULT 'enable',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE platforms;
DROP TABLE oses;
DROP TABLE manufacturers;
DROP TABLE isps;
DROP TABLE categories;
DROP TABLE browsers;

