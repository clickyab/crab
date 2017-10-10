
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO platforms (name,active,created_at,updated_at) VALUES ("desktop",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("mobile",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("tablet",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("other",1,"2017-09-11 08:02:33","2017-09-11 08:02:33");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM platforms;

