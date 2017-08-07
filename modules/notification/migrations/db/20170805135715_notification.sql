
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notification
(
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    message TEXT NOT NULL,
    type ENUM('sms', 'email', 'app') NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT notification_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
)
;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notification;
