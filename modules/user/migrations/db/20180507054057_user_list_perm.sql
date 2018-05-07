
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"user_list","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"user_list","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm IN ("user_list");