
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
TRUNCATE TABLE crab.role_permission;
ALTER TABLE role_permission MODIFY scope ENUM("self", "global", "superGlobal") NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


