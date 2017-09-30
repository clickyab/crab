-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE browsers;
CREATE TABLE browsers (
  'name'     VARCHAR(191) PRIMARY KEY,
  active     BOOL      NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX browsers_active
  ON browsers (active DESC);


DROP TABLE categories;
CREATE TABLE categories (
  'name'     VARCHAR(191) PRIMARY KEY,
  active     BOOL      NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX categories_active
  ON categories (active DESC);


DROP TABLE isps;
CREATE TABLE isps (
  'name'     VARCHAR(191) PRIMARY KEY,
  active     BOOL      NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX isps_active
  ON isps (active DESC);


DROP TABLE manufacturers;
CREATE TABLE manufacturers (
  'name'     VARCHAR(191) PRIMARY KEY,
  active     BOOL      NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX manufacturers_active
  ON manufacturers (active DESC);


DROP TABLE oses;
CREATE TABLE oses (
  'name'     VARCHAR(191) PRIMARY KEY,
  active     BOOL      NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX oses_active
  ON oses (active DESC);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE browsers;
CREATE TABLE browsers
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  name       VARCHAR(10)                         NOT NULL,
  active     TINYINT(1) DEFAULT '1'              NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT browsers_name_uindex
  UNIQUE (name)
);

DROP TABLE categories;
CREATE TABLE categories
(
  id          INT AUTO_INCREMENT
    PRIMARY KEY,
  name        VARCHAR(10)            NOT NULL,
  description VARCHAR(300)           NOT NULL,
  active      TINYINT(1) DEFAULT '1' NOT NULL,
  CONSTRAINT categoriesname_uindex
  UNIQUE (name)
);


DROP TABLE isps;
CREATE TABLE isps
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  name       VARCHAR(40)                             NOT NULL,
  active     TINYINT(1)                              NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     NOT NULL,
  updated_at TIMESTAMP DEFAULT '0000-00-00 00:00:00' NOT NULL,
  CONSTRAINT isps_name_uindex
  UNIQUE (name)
);

DROP TABLE manufacturers;
CREATE TABLE manufacturers
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  brand      VARCHAR(40)                             NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     NOT NULL,
  updated_at TIMESTAMP DEFAULT '0000-00-00 00:00:00' NOT NULL,
  active     TINYINT(1)                              NOT NULL,
  CONSTRAINT manufacturer_brand_uindex
  UNIQUE (brand)
);

DROP TABLE oses;
CREATE TABLE oses
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  name       VARCHAR(40)                             NOT NULL,
  active     TINYINT(1)                              NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     NOT NULL,
  updated_at TIMESTAMP DEFAULT '0000-00-00 00:00:00' NOT NULL,
  CONSTRAINT oses_name_uindex
  UNIQUE (name)
);

