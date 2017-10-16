
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope, created_at, updated_at) VALUES (2, 'get-campaign', 'self', '2017-10-16 12:01:57', '2017-10-16 12:01:58');
INSERT INTO role_permission (role_id, perm, scope, created_at, updated_at) VALUES (1, 'get-campaign', 'self', '2017-10-16 12:01:57', '2017-10-16 12:01:58');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm='get-campaign';

