
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO roles (name, description,level) VALUES ("SuperAdmin","SuperAdmin permission",5);
INSERT INTO roles (name, description,level) VALUES ("Owner","Owner permission",4);
INSERT INTO roles (name, description,level) VALUES ("Admin","Admin permission",3);
INSERT INTO roles (name, description,level) VALUES ("Advertiser","Advertiser permission",1);
INSERT INTO roles (name, description,level) VALUES ("Account","Account permission",2);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


