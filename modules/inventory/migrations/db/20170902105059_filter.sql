-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user_inventories (
  id         INT AUTO_INCREMENT PRIMARY KEY,
  created_at TIMESTAMP           NOT NULL,
  updated_at TIMESTAMP           NOT NULL,
  active     BOOL                NOT NULL,
  user_id    INT                 NOT NULL,
  label      VARCHAR(60)         NOT NULL  COMMENT 'user personal label',
  domains    TEXT                NOT NULL  COMMENT 'comma separated domains',
  kind       BOOL                NOT NULL  COMMENT 'whitelist = true, blacklist = false',
  publisher_type   ENUM ("web", "app") NOT NULL,
  CONSTRAINT inventory_user_label UNIQUE (user_id, label),
  INDEX inventory_active_update_user_id (user_id, active DESC, updated_at DESC)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_inventories;
