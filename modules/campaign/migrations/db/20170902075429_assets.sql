
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE campaign_manufacturer
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  brands TEXT
);

CREATE TABLE campaign_isp
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  isps TEXT
);

CREATE TABLE campaign_os
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  oss TEXT
);

CREATE TABLE campaign_category
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  oss TEXT
);

CREATE TABLE campaign_region
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  regions TEXT
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE campaign_manufacturer;
DROP TABLE campaign_isp;
DROP TABLE campaign_os;
DROP TABLE campaign_category;
DROP TABLE campaign_region;

