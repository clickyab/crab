
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO domains(`id`,`name`,`description`,`active`,`created_at`,`updated_at`) VALUES (1,'staging.crab.clickyab.ae','clickyab stuff',1 ,NOW(),NOW());
INSERT INTO domains(`id`,`name`,`description`,`active`,`created_at`,`updated_at`) VALUES (2,'127.0.0.1','local stuff',1 ,NOW(),NOW());
INSERT INTO domains(`id`,`name`,`description`,`active`,`created_at`,`updated_at`) VALUES (3,'localhost','local stuff', 1,NOW(),NOW());

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM domains WHERE id=1 OR name='staging.crab.clickyab.ae';
DELETE FROM domains WHERE id=2 OR name='127.0.0.1';
DELETE FROM domains WHERE id=3 OR name='localhost';