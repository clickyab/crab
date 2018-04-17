
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"create_bank_snap","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"create_bank_snap","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="create_bank_snap";

