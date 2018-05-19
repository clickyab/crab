
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DELETE FROM role_permission WHERE perm ="create_new_domain" AND role_id IN (3,4);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


