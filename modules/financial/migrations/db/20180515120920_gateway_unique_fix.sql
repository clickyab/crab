
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE gateways MODIFY name varchar(50) NOT NULL;
CREATE UNIQUE INDEX gateway_name_uindex ON gateways (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
