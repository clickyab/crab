
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"get_upload","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"get_upload","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM uploads WHERE perm="get_upload";

