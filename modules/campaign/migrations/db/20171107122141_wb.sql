
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_wlbl_presets DROP kind;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_wlbl_presets ADD COLUMN kind BOOL;

