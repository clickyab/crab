
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO role_permission (role_id, perm, scope) VALUES (1,"campaign_graph","self");
INSERT INTO role_permission (role_id, perm, scope) VALUES (2,"campaign_graph","self");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM role_permission WHERE perm ="campaign_graph";

