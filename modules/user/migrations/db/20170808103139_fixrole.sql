
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP INDEX roles_name_uindex ON roles;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ADD INDEX roles_name_uindex ON roles;


