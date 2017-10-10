
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE platforms
(
    name VARCHAR(15) PRIMARY KEY NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE paltforms;

