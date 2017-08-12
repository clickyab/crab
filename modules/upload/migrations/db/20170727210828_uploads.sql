
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table uploads
(
	id int auto_increment primary key,
	user_id int NOT NULL,
	path varchar(255) NOT null,
	mime VARCHAR(50) NOT NULL,
	size int NOT NULL,
	section VARCHAR(50) NOT NULL ,
	created_at timestamp default CURRENT_TIMESTAMP NOT null,
	constraint uploads_users_id_fk
	foreign key (user_id) references users (id)
)
;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE uploads;
