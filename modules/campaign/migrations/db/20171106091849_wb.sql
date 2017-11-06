
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE campaigns MODIFY white_black_type ENUM("all", "clickyab", "white", "black") NOT NULL DEFAULT "clickyab";

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE campaigns MODIFY white_black_type BOOL;


