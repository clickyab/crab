
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE UNIQUE INDEX roles_name_uindex ON roles (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


