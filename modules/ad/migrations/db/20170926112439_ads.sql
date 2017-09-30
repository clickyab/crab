
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ads
  ADD CONSTRAINT ads_campaign_id_fk FOREIGN KEY (campaign_id) REFERENCES campaigns (id);
  ALTER TABLE ads MODIFY src VARCHAR(191) NOT NULL;
  ALTER TABLE ads
  ADD CONSTRAINT ads_upload_id_fk FOREIGN KEY (src) REFERENCES uploads (id);
  CREATE UNIQUE INDEX ads_src_uindex ON ads (src);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ads
  DROP ads_campaign_id_fk;
  ALTER TABLE ads MODIFY src VARCHAR(511) NOT NULL;
  ALTER TABLE ads
  DROP ads_upload_id_fk;

