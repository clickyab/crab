
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaign_detail DROP PRIMARY KEY;
ALTER TABLE campaign_detail ADD PRIMARY KEY (campaign_id, hour_id, publisher_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaign_detail DROP PRIMARY KEY;
ALTER TABLE campaign_detail ADD PRIMARY KEY (campaign_id, hour_id);
