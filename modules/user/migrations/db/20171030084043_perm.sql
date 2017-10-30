
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (1,'get_banner','self',NOW(),NOW());
INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (2,'get_banner','self',NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE role_id IN (1,2) AND perm="get_banner";
