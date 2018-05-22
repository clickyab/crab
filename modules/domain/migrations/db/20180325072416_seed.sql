
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO domains (id,domain_base,description) VALUES (1,'staging.crab.clickyab.ae','clickyab staging');
INSERT INTO domains (id,domain_base,description) VALUES (2,'127.0.0.1','local stuff');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM domains WHERE id IN (1,2);

