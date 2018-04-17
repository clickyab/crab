
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DELETE FROM assets;
DELETE FROM creatives;
ALTER TABLE creatives ADD user_id INT(10) UNSIGNED NULL;
ALTER TABLE creatives
  ADD CONSTRAINT creatives_users_id_fk
FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE assets DROP FOREIGN KEY assets_creative_id_fk;
DROP INDEX unique_asset ON assets;
ALTER TABLE assets
  ADD CONSTRAINT assets_creatives_id_fk
FOREIGN KEY (creative_id) REFERENCES creatives (id);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


