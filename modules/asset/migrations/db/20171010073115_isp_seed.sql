
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO isps (name,active,kind,created_at,updated_at) VALUES ("irancell",1,"both","2017-09-11 08:02:33","2017-09-11 08:02:33"),
("hamraheaval",1,"both","2017-09-11 08:02:33","2017-09-11 08:02:33"),
("parsonline",1,"isp","2017-09-11 08:02:33","2017-09-11 08:02:33"),
("shatel",1,"isp","2017-09-11 08:02:33","2017-09-11 08:02:33"),
("talia",1,"cellular","2017-09-11 08:02:33","2017-09-11 08:02:33");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM isps;

