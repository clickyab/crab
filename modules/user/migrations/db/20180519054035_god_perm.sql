
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (7,"god","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (8,"god","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="god" AND scope="global" AND role_id IN (7,8);

