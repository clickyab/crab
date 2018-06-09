
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,'list_domain','superGlobal');
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,'get_domain','superGlobal');
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,'edit_domain','superGlobal');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE scope='superGlobal' AND perm IN ('list_domain', 'get_domain', 'edit_domain');