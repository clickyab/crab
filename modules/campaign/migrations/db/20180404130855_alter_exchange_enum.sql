
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `campaigns` CHANGE `exchange` `exchange` ENUM('clickyab','all_except_clickyab','all') DEFAULT 'clickyab';


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `campaigns` CHANGE `exchange` `exchange` TINYINT(1) NOT NULL;
