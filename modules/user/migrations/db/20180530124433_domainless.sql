
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD domain_less BOOL DEFAULT false  NOT NULL;
DROP TABLE crab.role_user;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


