
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `campaign_report_receivers` (
  `campaign_id` INT(10) UNSIGNED NOT NULL,
  `user_id` INT(10) UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

ALTER TABLE campaign_report_receivers
  ADD CONSTRAINT `campaign_report_receivers_campaigns_id_fk` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);
ALTER TABLE campaign_report_receivers
  ADD CONSTRAINT `campaign_report_receivers_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE campaign_report_receivers;