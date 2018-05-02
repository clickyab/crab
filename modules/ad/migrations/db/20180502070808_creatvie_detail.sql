
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `creative_detail` (
`campaign_id` int(10) unsigned NOT NULL,
`creative_id` int(10) unsigned NOT NULL,
`daily_id` int(10) unsigned NOT NULL,
`hour_id` int(10) unsigned NOT NULL,
`fake_imp` int(11) NOT NULL DEFAULT '0',
`fake_click` int(11) NOT NULL DEFAULT '0',
`imp` int(11) NOT NULL DEFAULT '0',
`click` int(11) NOT NULL DEFAULT '0',
`conv` int(11) NOT NULL DEFAULT '0',
`cpc` int(11) NOT NULL DEFAULT '0',
`cpm` int(11) NOT NULL DEFAULT '0',
`cpa` int(11) NOT NULL DEFAULT '0',
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`creative_id`,`hour_id`),
KEY `creative_detail_campaign_id_daily_id_index` (`campaign_id`,`daily_id`),
KEY `creative_detail_date_table_id_fk` (`daily_id`),
KEY `creative_detail_hour_table_id_fk` (`hour_id`),
CONSTRAINT `creative_detail_campaigns_id_fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`),
CONSTRAINT `creative_detail_date_table_id_fk` FOREIGN KEY (`daily_id`) REFERENCES `date_table` (`id`),
CONSTRAINT `creative_detail_hour_table_id_fk` FOREIGN KEY (`hour_id`) REFERENCES `hour_table` (`id`),
CONSTRAINT `creative_detail_creative_id_fk` FOREIGN KEY (`creative_id`) REFERENCES `creatives` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"campaign_creative","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"campaign_creative","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `creative_detail`;

DELETE  FROM `role_permission` WHERE perm='campaign_creative' AND scope="self" AND role_id IN(1,2);

