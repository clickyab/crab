
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE schedules MODIFY   h00         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h01         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h02         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h03         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h04         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h05         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h06         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h07         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h08         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h09         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h10         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h11         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h12         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h13         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h14         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h15         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h16         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h17         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h18         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h19         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h20         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h21         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h22         VARCHAR(100) NULL;
ALTER TABLE schedules MODIFY   h23         VARCHAR(100) NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE schedules MODIFY   h00         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h01         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h02         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h03         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h04         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h05         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h06         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h07         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h08         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h09         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h10         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h11         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h12         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h13         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h14         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h15         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h16         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h17         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h18         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h19         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h20         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h21         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h22         BOOL NOT NULL;
ALTER TABLE schedules MODIFY   h23         BOOL NOT NULL;

