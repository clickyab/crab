
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE consular_customer
(
  consulary_id INT NOT NULL,
  customer_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
  PRIMARY KEY (consulary_id,customer_id),
  CONSTRAINT cons_fk FOREIGN KEY (consulary_id) REFERENCES users (id),
  CONSTRAINT cost_fk FOREIGN KEY (customer_id) REFERENCES users (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE consular_customer;

