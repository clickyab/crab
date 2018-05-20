
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"log_campaign","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"log_campaign","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm IN ("log_campaign") AND role_id IN (3,4);


