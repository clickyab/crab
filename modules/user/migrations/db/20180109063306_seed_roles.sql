
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles(`id`,`name`,`description`,`domain_id`,`created_at`,`updated_at`) VALUES (1,'Advertiser','Advertiser role (default)',1,NOW(),NOW());
INSERT INTO roles(`id`,`name`,`description`,`domain_id`,`created_at`,`updated_at`) VALUES (2,'Advertiser','Advertiser role (default)',2,NOW(),NOW());
INSERT INTO roles(`id`,`name`,`description`,`domain_id`,`created_at`,`updated_at`) VALUES (3,'Advertiser','Advertiser role (default)',3,NOW(),NOW());

INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (1,1,'pub_list','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (2,2,'pub_list','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (3,3,'pub_list','self',NOW(),NOW());

INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (4,1,'assign_banner','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (5,2,'assign_banner','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (6,3,'assign_banner','self',NOW(),NOW());

INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (7,1,'inventory_list','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (8,2,'inventory_list','self',NOW(),NOW());
INSERT INTO role_permission(id, role_id, perm, scope, created_at, updated_at) VALUES (9,3,'inventory_list','self',NOW(),NOW());

INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (1,'get_banner','self',NOW(),NOW());
INSERT INTO role_permission(role_id, perm, scope, created_at, updated_at) VALUES (2,'get_banner','self',NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE role_id=1;
DELETE FROM role_permission WHERE role_id=2;
DELETE FROM role_permission WHERE role_id=3;
DELETE FROM role_permission WHERE role_id IN (1,2) AND perm="get_banner";
