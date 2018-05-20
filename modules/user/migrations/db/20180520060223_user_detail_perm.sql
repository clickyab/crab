
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"get_detail_user","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"get_detail_user","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="get_detail_user" AND scope="global" AND role_id IN (3,4);

