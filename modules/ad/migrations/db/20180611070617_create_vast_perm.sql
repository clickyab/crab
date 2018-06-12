
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (2, "create_vast_creative", "self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (3, "create_vast_creative", "self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4, "create_vast_creative", "self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (5, "create_vast_creative", "self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE  FROM `role_permission` WHERE perm='create_vast_creative' AND scope="self" AND role_id IN(2, 3, 4, 5);


