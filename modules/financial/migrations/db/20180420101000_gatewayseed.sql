
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO gateways (name, status, `is_default`, created_at, updated_at) VALUES ("saman","enable","yes","2018-04-17 14:28:52","2018-04-17 14:28:52");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM gateways WHERE name="saman";

