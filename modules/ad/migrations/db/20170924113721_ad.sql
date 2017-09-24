
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `ads` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`campaign_id` int(11) NOT NULL,
`src` varchar(511) NOT NULL,
`target` varchar(511) NOT NULL,
`width` int(11) NOT NULL,
`height` int(11) NOT NULL,
`status` enum('pending','accepted','rejected') NOT NULL,
`type` enum('banner','native','video') NOT NULL,
`mime` varchar(63) NOT NULL,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
PRIMARY KEY (`id`)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE ads;

