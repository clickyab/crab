
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (9,'un_removable','global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10,'un_removable','global');

INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'add_to_whitelabel_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'set_default_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'get_detail_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'get_detail_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'get_detail_domain', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'get_detail_domain', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'change_user_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'change_user_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'manual_change_cash', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'manual_change_cash', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'change_creatives_status', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'set_default_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'edit_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'add_to_whitelabel_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'user_list', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'user_list', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'edit_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'edit_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'impersonate_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'impersonate_user', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'add_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'add_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (9, 'edit_gateway', 'global');
INSERT INTO role_permission (role_id, perm, scope) VALUES (10, 'change_creatives_status', 'global');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE scope="global" AND role_id IN (9,10);