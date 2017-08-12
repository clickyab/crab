
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_personal DROP id;
ALTER TABLE user_personal ADD PRIMARY KEY (user_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_personal DROP PRIMARY KEY;
ALTER TABLE user_personal ADD COLUMN id PRIMARY;
ALTER TABLE user_personal DROP user_id;

