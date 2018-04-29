
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (id,name, description, domain_id) VALUES (1,'Advertiser','Advertiser role (default)',1);
INSERT INTO roles (id,name, description, domain_id) VALUES (2,'Advertiser','Advertiser role (default)',2);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE id IN (1,2);

