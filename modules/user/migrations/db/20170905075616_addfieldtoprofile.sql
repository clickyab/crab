
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_personal ADD COLUMN zip_code INT(11);
ALTER TABLE user_corporation ADD COLUMN zip_code INT(11);
ALTER TABLE user_personal ADD COLUMN national_id VARCHAR(30);
ALTER TABLE user_corporation ADD COLUMN national_id VARCHAR(30);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_personal DROP COLUMN zip_code;
ALTER TABLE user_corporation DROP COLUMN zip_code;
ALTER TABLE user_personal DROP COLUMN national_id;
ALTER TABLE user_corporation DROP COLUMN national_id;

