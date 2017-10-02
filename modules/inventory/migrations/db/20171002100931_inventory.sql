-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE publishers;
CREATE TABLE inventories (
  id        VARCHAR(191) PRIMARY KEY,
  create_at TIMESTAMP                               NOT NULL,
  update_at TIMESTAMP                               NOT NULL,
  active    BOOL                                    NOT NULL DEFAULT 1,
  name      VARCHAR(127)                            NOT NULL,
  domain    VARCHAR(127)                            NOT NULL,
  cat       TEXT                                    NULL,
  publisher VARCHAR(127)                            NOT NULL,
  kind      ENUM ('app', 'web')                     NOT NULL,
  status    ENUM ('pending', 'blocked', 'accepted') NOT NULL
);
CREATE INDEX inventories_domain ON inventories(domain(127));

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE inventories;
CREATE TABLE publishers
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  name VARCHAR(127) NOT NULL,
  domain VARCHAR(60) NOT NULL,
  pub_type ENUM("app", "web") NOT NULL,
  status ENUM("pending", "blocked", "accepted") NOT NULL,
  supplier VARCHAR(60) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT publishers_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
);


