
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (14,'un_removable','global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15,'un_removable','global');

INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'add_to_whitelabel_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'set_default_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'get_detail_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'get_detail_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'get_detail_domain', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'get_detail_domain', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'change_user_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'change_user_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'manual_change_cash', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'manual_change_cash', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'change_creatives_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'set_default_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'edit_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'add_to_whitelabel_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'user_list', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'user_list', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'edit_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'edit_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'impersonate_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'impersonate_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'add_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'add_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (14, 'edit_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (15, 'change_creatives_status', 'global');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE scope="global" AND role_id IN (14,15);