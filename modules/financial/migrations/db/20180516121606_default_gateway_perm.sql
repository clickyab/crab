
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"set_default_gateway","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"set_default_gateway","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="set_default_gateway" AND scope="global" AND role_id IN (3,4);

