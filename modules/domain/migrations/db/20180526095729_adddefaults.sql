
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE domains ADD min_total_budget int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_daily_budget int DEFAULT 0 NOT NULL;

ALTER TABLE domains ADD min_web_native_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_banner_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_vast_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_native_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_banner_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_vast_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_cpc int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_native_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_banner_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_vast_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_native_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_banner_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_vast_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_web_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD min_app_cpm int DEFAULT 0 NOT NULL;
ALTER TABLE domains ADD advantage int DEFAULT 0 NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE domains DROP min_total_budget;
ALTER TABLE domains DROP min_daily_budget;
ALTER TABLE domains DROP min_web_native_cpc;
ALTER TABLE domains DROP min_web_banner_cpc;
ALTER TABLE domains DROP min_web_vast_cpc;
ALTER TABLE domains DROP min_app_native_cpc;
ALTER TABLE domains DROP min_app_banner_cpc;
ALTER TABLE domains DROP min_app_vast_cpc;
ALTER TABLE domains DROP min_web_cpc;
ALTER TABLE domains DROP min_app_cpc;
ALTER TABLE domains DROP min_web_native_cpm;
ALTER TABLE domains DROP min_web_banner_cpm;
ALTER TABLE domains DROP min_web_vast_cpm;
ALTER TABLE domains DROP min_app_native_cpm;
ALTER TABLE domains DROP min_app_banner_cpm;
ALTER TABLE domains DROP min_app_vast_cpm;
ALTER TABLE domains DROP min_web_cpm;
ALTER TABLE domains DROP min_app_cpm;
ALTER TABLE domains DROP advantage;

