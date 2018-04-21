
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE online_payments MODIFY bank_status INT;
ALTER TABLE online_payments MODIFY ref_num VARCHAR(45);
ALTER TABLE online_payments MODIFY res_num VARCHAR(45) NOT NULL;
ALTER TABLE online_payments MODIFY status ENUM("init", "back_to_site", "finalized") NOT NULL;
ALTER TABLE gateways CHANGE `default` is_default ENUM("yes","no") NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


