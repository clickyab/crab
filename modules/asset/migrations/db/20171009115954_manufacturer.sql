
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE manufacturers DROP id;
ALTER TABLE manufacturers MODIFY brand VARCHAR(191) NOT NULL;
ALTER TABLE manufacturers ADD PRIMARY KEY (brand);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE manufacturers DROP brand;
ALTER TABLE manufacturers MODIFY brand VARCHAR(40) NOT NULL;
ALTER TABLE manufacturers ADD id INT NULL PRIMARY KEY;


