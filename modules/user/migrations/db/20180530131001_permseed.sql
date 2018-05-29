
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'change_user_status', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'change_user_status', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'edit_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'edit_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'get_detail_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'get_detail_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (1, 'impersonate_user', 'superGlobal');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'impersonate_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'impersonate_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (4, 'impersonate_user', 'self');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (5, 'impersonate_user', 'self');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'add_to_whitelabel_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'add_to_whitelabel_user', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (4, 'can_have_account', 'self');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'user_list', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'user_list', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (5, 'user_list', 'self');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (1, 'create_domain', 'superGlobal');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (1, 'change_domain_status', 'superGlobal');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'get_assign_admin_roles', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'get_assign_admin_roles', 'global');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


