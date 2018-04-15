
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE billings
  ADD CONSTRAINT `billings_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE billings
  ADD CONSTRAINT `billings_domain_id_fk` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`);

ALTER TABLE online_payments
  ADD CONSTRAINT `online_payments_domain_id_fk` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`);
  
ALTER TABLE online_payments
  ADD CONSTRAINT `online_payments_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE online_payments
  ADD CONSTRAINT `online_payments_gateway_id_fk` FOREIGN KEY (`gateway_id`) REFERENCES `gateways` (`id`);

ALTER TABLE bank_snaps
  ADD CONSTRAINT `bank_snaps_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE manual_cash_changes
  ADD CONSTRAINT `manual_cash_changes_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE manual_cash_changes
  ADD CONSTRAINT `manual_cash_changes_operator_id_fk` FOREIGN KEY (`operator_id`) REFERENCES `users` (`id`);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE billings
  DROP FOREIGN KEY `billings_user_id_fk`;

ALTER TABLE billings
  DROP FOREIGN KEY `billings_domain_id_fk`;

ALTER TABLE online_payments
  DROP FOREIGN KEY `online_payments_domain_id_fk`;

ALTER TABLE online_payments
  DROP FOREIGN KEY `online_payments_user_id_fk`;

ALTER TABLE online_payments
  DROP FOREIGN KEY `online_payments_gateway_id_fk`;

ALTER TABLE bank_snaps
  DROP FOREIGN KEY `bank_snaps_user_id_fk`;

ALTER TABLE manual_cash_changes
  DROP FOREIGN KEY `manual_cash_changes_user_id_fk`;

ALTER TABLE manual_cash_changes
  DROP FOREIGN KEY `manual_cash_changes_operator_id_fk`;