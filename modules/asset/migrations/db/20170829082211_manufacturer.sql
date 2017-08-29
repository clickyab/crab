
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE manufacturers
(
  id INT auto_increment
    PRIMARY KEY ,
  brand VARCHAR(40) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	active BOOL NOT NULL,
  CONSTRAINT manufacturer_brand_uindex
    UNIQUE (brand)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE manufacturers;
