
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE categories ADD status ENUM("enable", "disable") DEFAULT "enable" NOT NULL;
UPDATE categories SET categories.status="disable" WHERE name NOT IN (
  "IAB1",
  "IAB2",
  "IAB3",
  "IAB4",
  "IAB5",
  "IAB6",
  "IAB7",
  "IAB8",
  "IAB9",
  "IAB10",
  "IAB11",
  "IAB12",
  "IAB13",
  "IAB14",
  "IAB15",
  "IAB16",
  "IAB17",
  "IAB18",
  "IAB19",
  "IAB20",
  "IAB21",
  "IAB22",
  "IAB23",
  "IAB24",
  "IAB25",
  "IAB26");

CREATE TABLE category_model
(
  model_id int unsigned NOT NULL,
  category VARCHAR(15) NOT NULL,
  model ENUM("campaign", "publisher") NOT NULL
#   CONSTRAINT category_model_categories_name_fk FOREIGN KEY (category) REFERENCES categories (name)
);
CREATE UNIQUE INDEX category_model_model_id_category_model_uindex ON category_model (model_id, category, model);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE category_model;

