-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE campaigns (
  id                INT AUTO_INCREMENT PRIMARY KEY,
  created_at        TIMESTAMP                                    NOT NULL,
  updated_at        TIMESTAMP                                    NOT NULL,
  active            BOOL                                         NOT NULL DEFAULT 1,
  user_id           INT                                          NOT NULL,
  domain_id         INT                                          NOT NULL,
  kind              ENUM ("web", "app")                          NOT NULL,
  type              ENUM ("vast", "native", "banner")            NOT NULL,
  status            BOOL                                         NOT NULL
  COMMENT 'if campaign is active or not',
  start_at          TIMESTAMP                                    NOT NULL,
  end_at            TIMESTAMP,
  title             VARCHAR(255)                                 NOT NULL,
  budget            INT                                          NOT NULL,
  daily_limit       INT                                          NOT NULL,
  cost_type         ENUM ("cpm", "cpc", "cpa")                   NOT NULL,
  max_bid           INT                                          NOT NULL,
  notify_email      TEXT                                         NULL,
  progress           ENUM ("finalized", "inprogress") NOT NULL
  COMMENT 'determine if user is done with editing',
  white_black_id    INT,
  white_black_type  BOOL,
  white_black_value TEXT,
  CONSTRAINT campaign_domain_fk FOREIGN KEY (domain_id) REFERENCES domains (id),
  CONSTRAINT campaign_user_fk FOREIGN KEY (user_id) REFERENCES users (id),
  CONSTRAINT campaign_white_black_id_fk FOREIGN KEY (white_black_id) REFERENCES user_wlbl_presets (id)

);

CREATE TABLE campaign_attributes (
  campaign_id  INT  NOT NULL PRIMARY KEY,
  device       TEXT NULL,
  manufacturer TEXT NULL,
  os           TEXT NULL,
  browser      TEXT NULL,
  iab          TEXT NULL,
  region       TEXT NULL,
  cellular     TEXT NULL,
  isp          TEXT NULL,

  CONSTRAINT campaign_attribute_campaign__fk FOREIGN KEY (campaign_id
  ) REFERENCES campaigns (id
  )
);
ALTER TABLE schedules CHANGE `capmaign_id` `campaign_id` INT NOT NULL;

ALTER TABLE schedules
  ADD CONSTRAINT schedules_campaign_id_fk FOREIGN KEY (campaign_id) REFERENCES campaigns (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE schedules
  DROP schedules_campaign_id_fk;

ALTER TABLE schedules CHANGE `campaign_id` `capmaign_id`  INT NOT NULL;

DROP TABLE campaign_attributes;
DROP TABLE campaigns;


