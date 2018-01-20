
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ads ADD title VARCHAR(50) DEFAULT "" NOT NULL;
ALTER TABLE ads
  ALTER COLUMN title DROP DEFAULT;
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ads DROP title;

