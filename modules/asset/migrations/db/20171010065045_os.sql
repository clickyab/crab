
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE oses DROP id;
DROP INDEX oses_name_uindex ON oses;
ALTER TABLE oses ADD PRIMARY KEY (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE oses DROP PRIMARY KEY;
CREATE UNIQUE INDEX oses_name_uindex ON oses (name);
ALTER TABLE oses ADD id INT NOT NULL PRIMARY KEY;