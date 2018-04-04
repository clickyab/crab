
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"edit_campaign","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"edit_campaign","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"get_campaign","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"get_campaign","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"edit_attributes","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"edit_attributes","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"edit_budget","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"edit_budget","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm IN ("edit_campaign","get_campaign");

