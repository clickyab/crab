
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE creatives ADD name VARCHAR(64) NOT NULL;
ALTER TABLE creatives MODIFY max_bid INT UNSIGNED;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE creatives DROP name;
ALTER TABLE creatives MODIFY max_bid INT UNSIGNED NOT NULL;


