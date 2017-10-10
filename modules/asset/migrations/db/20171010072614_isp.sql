
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE isps DROP id;
DROP INDEX isps_name_uindex ON isps;
ALTER TABLE isps ADD PRIMARY KEY (name);
ALTER TABLE isps ADD kind ENUM("cellular", "isp", "both") NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE isps DROP COLUMN kind;
ALTER TABLE isps DROP PRIMARY KEY;
CREATE UNIQUE INDEX isps_name_uindex ON isps (name);
ALTER TABLE isps ADD id INT NOT NULL PRIMARY KEY;

