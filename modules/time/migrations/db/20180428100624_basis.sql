
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE date_table ADD COLUMN basis DATETIME;
UPDATE date_table set basis = STR_TO_DATE(concat(year,'-',month,'-',day), '%Y-%m-%e');
ALTER TABLE date_table MODIFY basis DATETIME NOT NULL ;


ALTER TABLE hour_table ADD COLUMN basis DATETIME;
UPDATE hour_table set basis = STR_TO_DATE(concat(year,'-',month,'-',day,'-',hour), '%Y-%m-%e-%k');
ALTER TABLE hour_table MODIFY basis DATETIME NOT NULL ;




-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE date_table DROP COLUMN basis;
ALTER TABLE hour_table DROP COLUMN basis;



