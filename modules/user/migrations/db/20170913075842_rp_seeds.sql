
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (1,1,'pub_list','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (2,2,'pub_list','self',NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM role_permission WHERE id=1;
DELETE FROM role_permission WHERE id=2;
