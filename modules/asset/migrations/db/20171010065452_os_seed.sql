
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO oses (name,active,created_at,updated_at) VALUES ("linux",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("windows",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("mac",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("android",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("ios",1,"2017-09-11 08:02:33","2017-09-11 08:02:33");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM oses;

