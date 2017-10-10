
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE categories DROP id;
DROP INDEX categoriesname_uindex ON categories;
ALTER TABLE categories MODIFY name VARCHAR(15) NOT NULL;
ALTER TABLE categories ADD PRIMARY KEY (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE categories DROP PRIMARY KEY;
ALTER TABLE categories ADD id INT NULL PRIMARY KEY;
CREATE UNIQUE INDEX categories_name_uindex ON categories (name);

