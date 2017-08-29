
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE oses(
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(40) NOT NULL,
  `active` BOOLEAN NOT NULL,
  created_at timestamp not null,
  updated_at timestamp not null
);
CREATE UNIQUE INDEX oses_name_uindex ON oses (`name`);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE oses;

