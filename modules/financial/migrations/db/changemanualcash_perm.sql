
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (3,"manual_change_cash","global");
INSERT INTO role_permission (role_id, perm, scope) VALUES (4,"manual_change_cash","global");
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm="manual_change_cash";