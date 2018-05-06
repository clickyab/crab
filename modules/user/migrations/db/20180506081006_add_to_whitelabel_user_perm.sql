
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (id,name, description, domain_id) VALUES (3,'Admin','Admin role (default)',1);
INSERT INTO roles (id,name, description, domain_id) VALUES (4,'Admin','Admin role (default)',2);

INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"add_to_whitelabel_user","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"add_to_whitelabel_user","global");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE  FROM `roles` WHERE id IN(3,4);
DELETE FROM role_permission WHERE perm IN ("add_to_whitelabel_user");


