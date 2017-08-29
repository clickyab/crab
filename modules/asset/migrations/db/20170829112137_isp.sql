
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE isps(
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(40) NOT NULL,
  `active` boolean NOT NULL,
  created_at timestamp not null,
  updated_at timestamp not null
);
CREATE UNIQUE INDEX isps_name_uindex ON isps (`name`);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE isps;

