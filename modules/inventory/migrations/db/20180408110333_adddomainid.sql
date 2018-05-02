
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE inventories ADD domain_id INT UNSIGNED NOT NULL;
ALTER TABLE inventories
  ADD CONSTRAINT inventories_domains_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

ALTER TABLE inventories DROP COLUMN publisher_type;
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


