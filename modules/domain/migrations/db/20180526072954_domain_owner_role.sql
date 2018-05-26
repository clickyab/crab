
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (id,name, description, domain_id) VALUES (9,'Owner','Domain Owner role',1);
INSERT INTO roles (id,name, description, domain_id) VALUES (10,'Owner','Domain Owner role',2);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE name="Owner" AND id IN (9,10);