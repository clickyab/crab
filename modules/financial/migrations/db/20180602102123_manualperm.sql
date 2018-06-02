
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (1, 'charge_owner', 'superGlobal');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (2, 'manual_change_cash', 'global');
INSERT INTO crab.role_permission (role_id, perm, scope) VALUES (3, 'manual_change_cash', 'global');


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


