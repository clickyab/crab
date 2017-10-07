
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE inventories MODIFY id INT AUTO_INCREMENT;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE inventories MODIFY id VARCHAR(191) NOT NULL;

