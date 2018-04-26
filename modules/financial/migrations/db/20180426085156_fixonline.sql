
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE online_payments ADD attr TEXT NULL;
ALTER TABLE online_payments MODIFY ref_num VARCHAR(127);
ALTER TABLE online_payments MODIFY res_num VARCHAR(127) NOT NULL;
ALTER TABLE online_payments MODIFY cid VARCHAR(127);
ALTER TABLE online_payments MODIFY trace_number VARCHAR(127);
ALTER TABLE online_payments ADD reason TEXT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE online_payments DROP COLUMN attr;
ALTER TABLE online_payments MODIFY ref_num VARCHAR(45);
ALTER TABLE online_payments MODIFY res_num VARCHAR(45) NOT NULL;
ALTER TABLE online_payments MODIFY cid VARCHAR(45);
ALTER TABLE online_payments MODIFY trace_number VARCHAR(45);
ALTER TABLE online_payments DROP COLUMN reason;

