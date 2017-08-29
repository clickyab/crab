-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE schedules (
  id          INT AUTO_INCREMENT PRIMARY KEY,
  capmaign_id INT       NOT NULL,
  updated_at  TIMESTAMP NOT NULL,
  h00         BOOL      NOT NULL,
  h01         BOOL      NOT NULL,
  h02         BOOL      NOT NULL,
  h03         BOOL      NOT NULL,
  h04         BOOL      NOT NULL,
  h05         BOOL      NOT NULL,
  h06         BOOL      NOT NULL,
  h07         BOOL      NOT NULL,
  h08         BOOL      NOT NULL,
  h09         BOOL      NOT NULL,
  h10         BOOL      NOT NULL,
  h11         BOOL      NOT NULL,
  h12         BOOL      NOT NULL,
  h13         BOOL      NOT NULL,
  h14         BOOL      NOT NULL,
  h15         BOOL      NOT NULL,
  h16         BOOL      NOT NULL,
  h17         BOOL      NOT NULL,
  h18         BOOL      NOT NULL,
  h19         BOOL      NOT NULL,
  h20         BOOL      NOT NULL,
  h21         BOOL      NOT NULL,
  h22         BOOL      NOT NULL,
  h23         BOOL      NOT NULL,
  CONSTRAINT schedules_campaign_id UNIQUE (capmaign_id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE schedules;

