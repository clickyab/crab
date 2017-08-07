
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table domains
(
	id int auto_increment
		primary key,
	name varchar(40) not null,
	description varchar(255) null,
	active enum('yes', 'no') default 'yes' not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	constraint domains_name_uindex
	unique (name)
)
;

create table domain_user
(
	domain_id int not null,
	user_id int not null,
	primary key (domain_id, user_id),
	constraint domain_user_domains_id_fk
	foreign key (domain_id) references domains (id),
	constraint domain_user_users_id_fk
	foreign key (user_id) references users (id)
)
;

create index domain_user_users_id_fk
	on domain_user (user_id)
;

alter table roles add constraint roles_domains_id_fk
foreign key (domain_id) references domains (id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE domain_user;
DROP TABLE domains;

