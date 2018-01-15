
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaign_detail ADD hour_id INT NOT NULL;
ALTER TABLE campaign_detail
  ADD CONSTRAINT campaign_detail_hour_table_id_fk
FOREIGN KEY (hour_id) REFERENCES hour_table (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaign_detail DROP FOREIGN KEY campaign_detail_hour_table_id_fk;
ALTER TABLE campaign_detail DROP hour_id;

