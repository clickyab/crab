
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `date_table` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`year` INT(10) UNSIGNED NOT NULL,
`month` INT(10) UNSIGNED NOT NULL,
`day` INT(10) UNSIGNED NOT NULL,
`j_year` INT(10) UNSIGNED NOT NULL,
`j_month` INT(10) UNSIGNED NOT NULL,
`j_day` INT(10) UNSIGNED NOT NULL,
`extra` TEXT NULL DEFAULT NULL,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `hour_table` (
`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
`year` INT(10) UNSIGNED NOT NULL,
`month` INT(10) UNSIGNED NOT NULL,
`day` INT(10) UNSIGNED NOT NULL,
`hour` INT(10) UNSIGNED NULL DEFAULT NULL,
`j_year` INT(10) UNSIGNED NOT NULL,
`j_month` INT(10) UNSIGNED NOT NULL,
`j_day` INT(10) UNSIGNED NOT NULL,
`date_id` INT(10) UNSIGNED NOT NULL,
PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE INDEX hour_time_date_table_id_fk
  ON hour_table (date_id);

ALTER TABLE hour_table
  ADD CONSTRAINT hour_time_date_table_id_fk
FOREIGN KEY (date_id) REFERENCES date_table (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE hour_table;
DROP TABLE date_table;

