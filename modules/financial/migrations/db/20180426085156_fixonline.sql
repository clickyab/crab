
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE online_payments ADD attr TEXT NULL;
ALTER TABLE online_payments MODIFY ref_num VARCHAR(127);
ALTER TABLE online_payments MODIFY res_num VARCHAR(127) NOT NULL;
ALTER TABLE online_payments MODIFY cid VARCHAR(127);
ALTER TABLE online_payments MODIFY trace_number VARCHAR(127);
ALTER TABLE online_payments ADD error_reason ENUM("v_invalid_request_data","v_params_illegal_characters","v_merchant_auth_failed","v_timeout_reached","v_empty_digital_code","v_params_too_long","v_invalid_return_amount","v_invalid_digital_code","v_params_too_short","v_negative_amount","v_amount_not_match","v_transaction_not_found","v_invalid_amount","v_internal_bank","v_multi_amount","v_invalid_ip_or_pass","v_not_supported","p_cancel_by_user","p_amount_not_match","p_early_verify","p_invalid_card_num","p_invalid_card_owner","p_card_expired","p_wrong_pass_3_times","p_wrong_pass","p_amount_exceed","p_pin_pan_error","p_response_timeout","p_invalid_cvv_or_expdate","p_no_sufficient_funds","p_card_issuer_is_down","p_bank_error","p_not_supported","hash_mismatch","merchant_mismatch","verify_request_failed","verify_response_failed","amount_mismatch");


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE online_payments DROP COLUMN attr;
ALTER TABLE online_payments MODIFY ref_num VARCHAR(45);
ALTER TABLE online_payments MODIFY res_num VARCHAR(45) NOT NULL;
ALTER TABLE online_payments MODIFY cid VARCHAR(45);
ALTER TABLE online_payments MODIFY trace_number VARCHAR(45);
ALTER TABLE online_payments DROP COLUMN reason;

