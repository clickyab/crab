
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (id,name, description, domain_id) VALUES (7,'SuperAdmin','Super Admin role',1);
INSERT INTO roles (id,name, description, domain_id) VALUES (8,'SuperAdmin','Super Admin role',2);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE id IN (7,8);
