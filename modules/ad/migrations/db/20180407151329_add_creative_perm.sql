
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"add_creative","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"add_creative","self");

INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"edit_creative","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"edit_creative","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE  FROM `role_permission` WHERE perm='add_creative' AND scope="self" AND role_id IN(1,2);
DELETE  FROM `role_permission` WHERE perm='edit_creative' AND scope="self" AND role_id IN(1,2);
