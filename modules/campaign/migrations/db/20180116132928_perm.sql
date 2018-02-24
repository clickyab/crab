
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (1,'campaign_graph','self',NOW(),NOW());
INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (2,'campaign_graph','self',NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="campaign_graph";