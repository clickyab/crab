
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE browsers DROP id;
DROP INDEX browsers_name_uindex ON browsers;
ALTER TABLE browsers MODIFY name VARCHAR(60) NOT NULL;
ALTER TABLE browsers ADD PRIMARY KEY (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE browsers DROP PRIMARY KEY;
ALTER TABLE browsers MODIFY name VARCHAR(10) NOT NULL;
ALTER TABLE browsers ADD id INT NULL PRIMARY KEY;

