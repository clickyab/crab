
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users MODIFY avatar VARCHAR(191);
CREATE UNIQUE INDEX users_avatar_uindex ON users (avatar);
ALTER TABLE users
ADD CONSTRAINT users_uploads_id_fk
FOREIGN KEY (avatar) REFERENCES uploads (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP FOREIGN KEY users_uploads_id_fk;
DROP INDEX users_avatar_uindex ON users;
ALTER TABLE users MODIFY avatar VARCHAR(255);

