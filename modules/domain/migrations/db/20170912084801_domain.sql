
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE domains DROP active;
ALTER TABLE domains ADD active BOOL DEFAULT 1 NOT NULL ;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE domains DROP active;
ALTER TABLE domains ADD active enum('yes', 'no') default 'yes' NOT NULL ;


