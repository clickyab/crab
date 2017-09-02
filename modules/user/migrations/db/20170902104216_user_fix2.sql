
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_personal ADD province_id int null;
ALTER TABLE user_personal ADD country_id int null;

ALTER TABLE user_corporation ADD province_id int null;
ALTER TABLE user_corporation ADD country_id int null;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_personal DROP province_id;
ALTER TABLE user_personal DROP country_id;

ALTER TABLE user_corporation DROP province_id;
ALTER TABLE user_corporation DROP country_id ;

