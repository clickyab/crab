
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `campaigns` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` INT(10) UNSIGNED NOT NULL,
`domain_id` INT(10) UNSIGNED NOT NULL,
`kind` ENUM('web', 'app') NOT NULL,
`status` ENUM('start', 'pause') NOT NULL COMMENT 'if campaign is active or not',
`start_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`end_at` DATETIME NULL DEFAULT NULL,
`title` VARCHAR(255) NOT NULL,
`tld` VARCHAR(255) NULL DEFAULT NULL COMMENT 'top level domain',
`total_budget` INT(10) UNSIGNED NOT NULL,
`daily_budget` INT(10) UNSIGNED NOT NULL,
`today_spend` INT(10) UNSIGNED NOT NULL DEFAULT 0,
`total_spend` INT(10) UNSIGNED NOT NULL DEFAULT 0,
`max_bid` INT(10) NOT NULL DEFAULT 0,
`strategy` ENUM('cpc', 'cpa', 'cpm') NOT NULL,
`progress` ENUM('inprogress', 'finalized') NOT NULL COMMENT 'determine if user is done with editing',
`inventory_id` INT(10) UNSIGNED NULL DEFAULT NULL,
`inventory_type` ENUM('black_list', 'white_list') NULL DEFAULT NULL,
`inventory_domains` TEXT NULL DEFAULT NULL,
`exchange` TINYINT(1) NOT NULL,
`archived_at` DATETIME NULL DEFAULT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `campaign_detail` (
  `campaign_id` INT(10) UNSIGNED NOT NULL,
  `daily_id` INT(10) UNSIGNED NOT NULL,
  `hour_id` INT(10) UNSIGNED NOT NULL,
  `fake_imp` INT(11) NOT NULL DEFAULT '0',
  `fake_click` INT(11) NOT NULL DEFAULT '0',
  `imp` INT(11) NOT NULL DEFAULT '0',
  `click` INT(11) NOT NULL DEFAULT '0',
  `conv` INT(11) NOT NULL DEFAULT '0',
  `cpc` INT(11) NOT NULL DEFAULT '0',
  `cpm` INT(11) NOT NULL DEFAULT '0',
  `cpa` INT(11) NOT NULL DEFAULT '0',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`campaign_id`, `hour_id`),
  INDEX `campaign_detail_campaign_id_daily_id_index` (`campaign_id` ASC, `daily_id` ASC))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `campaign_notify` (
  `campaign_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `schedules` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `campaign_id` INT(10) UNSIGNED NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `h00` VARCHAR(100) NULL DEFAULT NULL,
  `h01` VARCHAR(100) NULL DEFAULT NULL,
  `h02` VARCHAR(100) NULL DEFAULT NULL,
  `h03` VARCHAR(100) NULL DEFAULT NULL,
  `h04` VARCHAR(100) NULL DEFAULT NULL,
  `h05` VARCHAR(100) NULL DEFAULT NULL,
  `h06` VARCHAR(100) NULL DEFAULT NULL,
  `h07` VARCHAR(100) NULL DEFAULT NULL,
  `h08` VARCHAR(100) NULL DEFAULT NULL,
  `h09` VARCHAR(100) NULL DEFAULT NULL,
  `h10` VARCHAR(100) NULL DEFAULT NULL,
  `h11` VARCHAR(100) NULL DEFAULT NULL,
  `h12` VARCHAR(100) NULL DEFAULT NULL,
  `h13` VARCHAR(100) NULL DEFAULT NULL,
  `h14` VARCHAR(100) NULL DEFAULT NULL,
  `h15` VARCHAR(100) NULL DEFAULT NULL,
  `h16` VARCHAR(100) NULL DEFAULT NULL,
  `h17` VARCHAR(100) NULL DEFAULT NULL,
  `h18` VARCHAR(100) NULL DEFAULT NULL,
  `h19` VARCHAR(100) NULL DEFAULT NULL,
  `h20` VARCHAR(100) NULL DEFAULT NULL,
  `h21` VARCHAR(100) NULL DEFAULT NULL,
  `h22` VARCHAR(100) NULL DEFAULT NULL,
  `h23` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `campaign_attributes` (
  `campaign_id` INT(10) UNSIGNED NOT NULL,
  `device` TEXT NULL DEFAULT NULL,
  `manufacturer` TEXT NULL DEFAULT NULL,
  `os` TEXT NULL DEFAULT NULL,
  `browser` TEXT NULL DEFAULT NULL,
  `iab` TEXT NULL DEFAULT NULL,
  `region` TEXT NULL DEFAULT NULL,
  `cellular` TEXT NULL DEFAULT NULL,
  `isp` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`campaign_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8;

ALTER TABLE campaigns
  ADD CONSTRAINT campaign_domain_id_fk
FOREIGN KEY (domain_id) REFERENCES domains (id);

ALTER TABLE campaigns
  ADD CONSTRAINT campaign_user_fk
FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE campaigns
  ADD CONSTRAINT `campaign_domain_fk` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`);
ALTER TABLE campaigns
  ADD CONSTRAINT `campaign_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE campaign_notify
  ADD CONSTRAINT `campaign_notify_campaigns_id_fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);
ALTER TABLE campaign_notify
  ADD CONSTRAINT `campaign_notify_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE campaign_attributes
  ADD CONSTRAINT `campaign_attribute_campaign__fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);

ALTER TABLE campaign_detail
  ADD CONSTRAINT `campaign_detail_campaigns_id_fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);
ALTER TABLE campaign_detail
  ADD CONSTRAINT `campaign_detail_date_table_id_fk` FOREIGN KEY (`daily_id`) REFERENCES `date_table` (`id`);
ALTER TABLE campaign_detail
  ADD CONSTRAINT `campaign_detail_hour_table_id_fk` FOREIGN KEY (`hour_id`) REFERENCES `hour_table` (`id`);

ALTER TABLE schedules
  ADD CONSTRAINT schedules_campaign_id_fk
FOREIGN KEY (campaign_id) REFERENCES campaigns (id);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE campaign_attributes;
DROP TABLE schedules;
DROP TABLE campaign_notify;
DROP TABLE campaign_detail;
DROP TABLE campaigns;

