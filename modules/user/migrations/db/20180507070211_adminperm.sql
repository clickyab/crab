
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (id,name, description, domain_id) VALUES (5,'Agency','Agency role',1);
INSERT INTO roles (id,name, description, domain_id) VALUES (6,'Agency','Agency role',2);

INSERT INTO role_permission (role_id, perm, scope) VALUES (5,"can_manage","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (6,"can_manage","self");

INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"edit_user","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"edit_user","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="edit_user" AND role_id IN (3,4);
DELETE FROM role_permission WHERE perm="can_manage" AND role_id IN (5,6);
DELETE FROM roles WHERE id IN (5,6);


