
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO domains(`id`,`name`,`description`,`active`,`created_at`,`updated_at`) VALUES (1,'staging.crab.clickyab.ae','clickyab stuff','yes',NOW(),NOW());
INSERT INTO domains(`id`,`name`,`description`,`active`,`created_at`,`updated_at`) VALUES (2,'127.0.0.1','local stuff','yes',NOW(),NOW());
INSERT INTO roles(`id`,`name`,`description`,`domain_id`,`created_at`,`updated_at`) VALUES (1,'Advertiser','Advertiser role (default)',1,NOW(),NOW());
INSERT INTO roles(`id`,`name`,`description`,`domain_id`,`created_at`,`updated_at`) VALUES (2,'Advertiser','Advertiser role (default)',2,NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM roles WHERE id=1;
DELETE FROM domains WHERE id=1;
DELETE FROM domains WHERE id=2;

