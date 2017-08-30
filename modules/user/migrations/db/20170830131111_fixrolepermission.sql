
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE role_permission ADD UNIQUE `unique_group`(role_id, perm, scope);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE role_permission DROP COLUMN 'unique_group';

