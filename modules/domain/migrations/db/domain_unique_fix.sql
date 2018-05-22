
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE UNIQUE INDEX domains_name_uindex ON domains (name);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

