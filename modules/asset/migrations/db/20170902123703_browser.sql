
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE browsers
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(10) NOT NULL,
  `active` boolean NOT NULL DEFAULT true,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL
);
CREATE UNIQUE INDEX browsers_name_uindex ON browsers (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE browsers;

