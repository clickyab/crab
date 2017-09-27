
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DELETE FROM user_wlbl_presets;
ALTER TABLE user_wlbl_presets ADD COLUMN domain_id INT NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_wlbl_presets DROP COLUMN domain_id;

