
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE campaign_document
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  os TEXT,
  browser TEXT,
  brand TEXT,
  category TEXT,
  isp TEXT,
  created_at timestamp not null,
  updated_at timestamp not null
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE campaign_document;
