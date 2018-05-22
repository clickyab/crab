
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE UNIQUE INDEX domains_domain_base_uindex ON domains (domain_base);
ALTER TABLE domains ADD title varchar(127) NULL;
ALTER TABLE domains ADD logo varchar(127) NULL;
ALTER TABLE domains ADD theme varchar(31) NULL;

ALTER TABLE domains MODIFY title varchar(127) NOT NULL;
ALTER TABLE domains MODIFY theme varchar(31) NOT NULL;

UPDATE domains SET title="clickyab",theme="red";

ALTER TABLE domains MODIFY logo varchar(191);
ALTER TABLE domains
  ADD CONSTRAINT domains_uploads_id_fk
FOREIGN KEY (logo) REFERENCES uploads (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE domains CHANGE domain_base name varchar(254) NOT NULL;
ALTER TABLE domains DROP title;
ALTER TABLE domains DROP logo;
ALTER TABLE domains DROP theme;

