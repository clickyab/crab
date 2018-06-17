
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS browsers;
CREATE TABLE `browsers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `desktop` tinyint(1) NOT NULL DEFAULT '1',
  `mobile` tinyint(1) NOT NULL DEFAULT '1',
  `tablet` tinyint(1) DEFAULT '1',
  `other` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `browsers_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (1,"Google Chrome",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (2,"Firefox",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (3,"Safari",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (4,"Internet Explorer",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (5,"Edge",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (6,"Opera",true ,false ,false ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (7,"Maxthon",true ,true ,true ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (8,"iOS Safari",false ,true ,true ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (9,"Android Browser",false ,true ,true ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (10,"Chrome For Android",false ,true ,true ,true);
INSERT INTO browsers (id, name, desktop, mobile, tablet, other) VALUES (11,"Opera Mini",false ,true ,true ,true);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE browsers;

