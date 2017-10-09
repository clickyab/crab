
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO browsers (name,active,created_At,updated_at) VALUES ("Mozilla Firefox",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("Google Chrome",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("Opera",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("Microsoft Edge",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("Internet Explorer",1,"2017-09-11 08:02:33","2017-09-11 08:02:33"),
 ("Safari",1,"2017-09-11 08:02:33","2017-09-11 08:02:33");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM browsers;

