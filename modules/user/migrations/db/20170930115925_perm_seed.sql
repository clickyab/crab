
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (3,1,'assign_banner','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (4,2,'assign_banner','self',NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM role_permission WHERE id=3;
DELETE FROM role_permission WHERE id=4;
