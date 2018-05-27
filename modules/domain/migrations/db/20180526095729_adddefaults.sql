
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE domains ADD total_budget int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD daily_budget int DEFAULT 0 NOT NULL;

ALTER TABLE domains ADD web_native_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_banner_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_vast_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_native_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_banner_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_vast_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_native_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_banner_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_vast_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_native_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_banner_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_vast_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD web_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD app_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD advantage int DEFAULT 0 NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE domains DROP total_budget;
ALTER TABLE domains DROP daily_budget;
ALTER TABLE domains DROP web_native_cpc;
ALTER TABLE domains DROP web_banner_cpc;
ALTER TABLE domains DROP web_vast_cpc;
ALTER TABLE domains DROP app_native_cpc;
ALTER TABLE domains DROP app_banner_cpc;
ALTER TABLE domains DROP app_vast_cpc;
ALTER TABLE domains DROP web_cpc;
ALTER TABLE domains DROP app_cpc;
ALTER TABLE domains DROP web_native_cpm;
ALTER TABLE domains DROP web_banner_cpm;
ALTER TABLE domains DROP web_vast_cpm;
ALTER TABLE domains DROP app_native_cpm;
ALTER TABLE domains DROP app_banner_cpm;
ALTER TABLE domains DROP app_vast_cpm;
ALTER TABLE domains DROP web_cpm;
ALTER TABLE domains DROP app_cpm;
ALTER TABLE domains DROP advantage;

