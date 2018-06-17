
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS oses;

CREATE TABLE `oses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL,
  `status` enum('disable','enable') NOT NULL DEFAULT 'enable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `app` tinyint(1) NOT NULL DEFAULT '1',
  `web` tinyint(1) NOT NULL DEFAULT '1',
  `other` tinyint(1) NOT NULL DEFAULT '1',
  `tablet` tinyint(1) NOT NULL DEFAULT '1',
  `mobile` tinyint(1) NOT NULL DEFAULT '1',
  `desktop` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `oses_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;


CREATE TABLE `os_browser` (
  `browser_id` int(11) NOT NULL,
  `os_id` int(11) NOT NULL,
  PRIMARY KEY (`browser_id`,`os_id`),
  KEY `os_browser_oses_id_fk` (`os_id`),
  CONSTRAINT `os_browser_browsers_id_fk` FOREIGN KEY (`browser_id`) REFERENCES `browsers` (`id`),
  CONSTRAINT `os_browser_oses_id_fk` FOREIGN KEY (`os_id`) REFERENCES `oses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (1,"Windows",false ,true ,true,false ,false ,true);
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (2,"OSX (Mac)",false ,true ,true,false ,false ,true);
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (3,"Linux/Unix/BSD",false ,true ,true,false ,false ,true);
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (4,"Android",true ,false ,true,true ,true ,false );
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (5,"WindowsPhone",true ,false ,true,true ,true ,false );
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (6,"iOS",true ,false ,true,true ,true ,false );
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (7,"Bada",true ,false ,true,true ,true ,true);
INSERT INTO oses (id, name, app, web, other, tablet, mobile, desktop) VALUES (8,"Others",true ,true ,true,true ,true ,true);

INSERT INTO os_browser (os_id,browser_id) VALUES (1,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (1,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (1,4);
INSERT INTO os_browser (os_id,browser_id) VALUES (1,5);
INSERT INTO os_browser (os_id,browser_id) VALUES (1,6);
INSERT INTO os_browser (os_id,browser_id) VALUES (1,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (2,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (2,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (2,3);
INSERT INTO os_browser (os_id,browser_id) VALUES (2,6);
INSERT INTO os_browser (os_id,browser_id) VALUES (2,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (3,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (3,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (3,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (4,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (4,9);
INSERT INTO os_browser (os_id,browser_id) VALUES (4,10);
INSERT INTO os_browser (os_id,browser_id) VALUES (4,11);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,4);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,5);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,6);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (5,11);
INSERT INTO os_browser (os_id,browser_id) VALUES (6,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (6,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (6,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (6,8);
INSERT INTO os_browser (os_id,browser_id) VALUES (6,11);
INSERT INTO os_browser (os_id,browser_id) VALUES (7,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (7,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (7,6);
INSERT INTO os_browser (os_id,browser_id) VALUES (7,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,1);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,2);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,3);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,4);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,5);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,6);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,7);
INSERT INTO os_browser (os_id,browser_id) VALUES (8,11);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE oses;
DROP TABLE os_browser;

