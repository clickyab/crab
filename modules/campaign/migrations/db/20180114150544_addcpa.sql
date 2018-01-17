
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE campaign_detail;
CREATE TABLE `campaign_detail` (
`campaign_id` int(11) NOT NULL,
`daily_id` int(11) NOT NULL,
`fake_imp` int(11) NOT NULL DEFAULT '0',
`fake_click` int(11) NOT NULL DEFAULT '0',
`imp` int(11) NOT NULL DEFAULT '0',
`click` int(11) NOT NULL DEFAULT '0',
`conv` int(11) NOT NULL DEFAULT '0',
`cpc` int(11) NOT NULL DEFAULT '0',
`cpm` int(11) NOT NULL DEFAULT '0',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`cpa` int(11) NOT NULL DEFAULT '0',
`hour_id` int(11) NOT NULL,
PRIMARY KEY (`campaign_id`,`hour_id`),
KEY `campaign_detail_date_table_id_fk` (`daily_id`),
KEY `campaign_detail_campaign_id_daily_id_index` (`campaign_id`,`daily_id`),
KEY `campaign_detail_hour_table_id_fk` (`hour_id`),
CONSTRAINT `campaign_detail_campaigns_id_fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`),
CONSTRAINT `campaign_detail_date_table_id_fk` FOREIGN KEY (`daily_id`) REFERENCES `date_table` (`id`),
CONSTRAINT `campaign_detail_hour_table_id_fk` FOREIGN KEY (`hour_id`) REFERENCES `hour_table` (`id`)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE campaign_detail;