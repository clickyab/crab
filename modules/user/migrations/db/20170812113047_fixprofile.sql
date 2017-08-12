
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_corporation DROP id;
ALTER TABLE user_corporation ADD PRIMARY KEY (user_id);
ALTER TABLE user_corporation ADD COLUMN last_name VARCHAR(40) NOT NULL;
ALTER TABLE user_corporation ADD COLUMN first_name VARCHAR(40) NOT NULL;
ALTER TABLE user_personal MODIFY COLUMN last_name VARCHAR(40) NOT NULL;
ALTER TABLE user_personal MODIFY COLUMN first_name VARCHAR(40) NOT NULL;
ALTER TABLE user_corporation MODIFY COLUMN name VARCHAR(50) NOT NULL;
ALTER TABLE user_personal MODIFY gender ENUM('male', 'female', 'not_specified');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_corporation DROP PRIMARY KEY;
ALTER TABLE user_corporation ADD COLUMN id PRIMARY;
ALTER TABLE user_corporation DROP user_id;

