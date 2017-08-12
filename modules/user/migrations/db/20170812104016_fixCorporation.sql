
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_corporation DROP id;
ALTER TABLE user_corporation ADD PRIMARY KEY (user_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
;
